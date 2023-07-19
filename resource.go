//go:build darwin
// +build darwin

package mtl

/*
#include "resource.h"
*/
import "C"
import "unsafe"

// Buffer is a memory allocation for storing unformatted data
// that is accessible to the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtlbuffer
type Buffer struct {
	buffer unsafe.Pointer
}

// resource implements the Resource interface.
func (b *Buffer) resource() unsafe.Pointer { return b.buffer }

// Contents returns a pointer to the contents of the buffer.
func (b *Buffer) Contents() unsafe.Pointer {
	return C.Buffer_Contents(b.buffer)
}

// NewBufferWithLength creates a new buffer with the specified length.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/1433375-newbufferwithlength
func (d Device) NewBufferWithLength(length uintptr, opt ResourceOptions) Buffer {
	b := C.Device_NewBufferWithLength(d.device, C.size_t(length), C.uint16_t(opt))
	return Buffer{buffer: b}
}

// NewBufferWithBytes creates a new buffer of a given length and initializes its contents by copying existing data into it.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/1433429-newbufferwithbytes
func (d Device) NewBufferWithBytes(bytes unsafe.Pointer, length uintptr, opt ResourceOptions) Buffer {
	b := C.Device_NewBufferWithBytes(d.device, bytes, C.size_t(length), C.uint16_t(opt))
	return Buffer{buffer: b}
}

// TextureDescriptor configures new Texture objects.
//
// Reference: https://developer.apple.com/documentation/metal/mtltexturedescriptor
type TextureDescriptor struct {
	PixelFormat PixelFormat
	Width       uint
	Height      uint
	StorageMode StorageMode
}

// Texture is a memory allocation for storing formatted
// image data that is accessible to the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtltexture
type Texture struct {
	texture unsafe.Pointer

	// Width is the width of the texture image for the base level mipmap, in pixels.
	Width uint

	// Height is the height of the texture image for the base level mipmap, in pixels.
	Height uint
}

// NewTextureWithDescriptor creates a new texture with the provided descriptor using the device.
func (d Device) NewTextureWithDescriptor(td TextureDescriptor) Texture {
	descriptor := C.struct_TextureDescriptor{
		PixelFormat: C.uint16_t(td.PixelFormat),
		Width:       C.uint_t(td.Width),
		Height:      C.uint_t(td.Height),
		StorageMode: C.uint8_t(td.StorageMode),
	}

	return Texture{
		texture: C.Device_NewTextureWithDescriptor(d.device, descriptor),
		Width:   td.Width,
		Height:  td.Height,
	}
}

// resource implements the Resource interface.
func (t Texture) resource() unsafe.Pointer { return t.texture }

// ReplaceRegion copies a block of pixels into a section of texture slice 0.
//
// Reference: https://developer.apple.com/documentation/metal/mtltexture/1515464-replaceregion
func (t Texture) ReplaceRegion(region Region, level int, pixelBytes *byte, bytesPerRow uintptr) {
	r := C.struct_Region{
		Origin: C.struct_Origin{
			X: C.uint_t(region.Origin.X),
			Y: C.uint_t(region.Origin.Y),
			Z: C.uint_t(region.Origin.Z),
		},
		Size: C.struct_Size{
			Width:  C.uint_t(region.Size.Width),
			Height: C.uint_t(region.Size.Height),
			Depth:  C.uint_t(region.Size.Depth),
		},
	}
	C.Texture_ReplaceRegion(t.texture, r, C.uint_t(level), unsafe.Pointer(pixelBytes), C.size_t(bytesPerRow))
}

// GetBytes copies a block of pixels from the storage allocation of texture
// slice zero into system memory at a specified address.
//
// Reference: https://developer.apple.com/documentation/metal/mtltexture/1515751-getbytes
func (t Texture) GetBytes(pixelBytes *byte, bytesPerRow uintptr, region Region, level int) {
	r := C.struct_Region{
		Origin: C.struct_Origin{
			X: C.uint_t(region.Origin.X),
			Y: C.uint_t(region.Origin.Y),
			Z: C.uint_t(region.Origin.Z),
		},
		Size: C.struct_Size{
			Width:  C.uint_t(region.Size.Width),
			Height: C.uint_t(region.Size.Height),
			Depth:  C.uint_t(region.Size.Depth),
		},
	}
	C.Texture_GetBytes(t.texture, unsafe.Pointer(pixelBytes), C.size_t(bytesPerRow), r, C.uint_t(level))
}
