//go:build darwin
// +build darwin

package mtl

/*
#include "command_encoder.h"
*/
import "C"
import (
	"unsafe"
)

// CommandEncoder is an encoder that writes sequential GPU commands
// into a command buffer.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandencoder
type CommandEncoder struct {
	commandEncoder unsafe.Pointer
}

// EndEncoding declares that all command generation from this encoder is completed.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandencoder/1458038-endencoding
func (ce CommandEncoder) EndEncoding() {
	C.CommandEncoder_EndEncoding(ce.commandEncoder)
}

// ComputeCommandEncoder is an object for encoding commands in a compute pass.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder
type ComputeCommandEncoder struct {
	CommandEncoder
}

// SetComputePipelineState sets the current compute pipeline state object for the compute command encoder.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443140-setcomputepipelinestate
func (cce ComputeCommandEncoder) SetComputePipelineState(cps ComputePipelineState) {
	C.ComputeCommandEncoder_SetComputePipelineState(cce.commandEncoder, cps.computePipelineState)
}

// SetBuffer sets a buffer for the compute function at a specified index.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443126-setbuffer
func (cce ComputeCommandEncoder) SetBuffer(buf Buffer, offset, index int) {
	C.ComputeCommandEncoder_SetBuffer(cce.commandEncoder, buf.buffer, C.uint_t(offset), C.uint_t(index))
}

// DispatchThreads encodes a compute command using an arbitrarily sized grid.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2866532-dispatchthreads
func (cce ComputeCommandEncoder) DispatchThreads(gridSize, threadgroupSize Size) {
	gs := C.struct_Size{
		Width:  C.uint_t(gridSize.Width),
		Height: C.uint_t(gridSize.Height),
		Depth:  C.uint_t(gridSize.Depth),
	}

	tgs := C.struct_Size{
		Width:  C.uint_t(threadgroupSize.Width),
		Height: C.uint_t(threadgroupSize.Height),
		Depth:  C.uint_t(threadgroupSize.Depth),
	}

	C.ComputeCommandEncoder_DispatchThreads(cce.commandEncoder, gs, tgs)
}

// RenderCommandEncoder is an encoder that specifies graphics-rendering commands
// and executes graphics functions.
//
// Reference: https://developer.apple.com/documentation/metal/mtlrendercommandencoder.
type RenderCommandEncoder struct {
	CommandEncoder
}

// SetRenderPipelineState sets the current render pipeline state object.
//
// Reference: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515811-setrenderpipelinestate.
func (rce RenderCommandEncoder) SetRenderPipelineState(rps RenderPipelineState) {
	C.RenderCommandEncoder_SetRenderPipelineState(rce.commandEncoder, rps.renderPipelineState)
}

// SetVertexBuffer sets a buffer for the vertex shader function at an index
// in the buffer argument table with an offset that specifies the start of the data.
//
// Reference: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515829-setvertexbuffer.
func (rce RenderCommandEncoder) SetVertexBuffer(buf Buffer, offset, index int) {
	C.RenderCommandEncoder_SetVertexBuffer(rce.commandEncoder, buf.buffer, C.uint_t(offset), C.uint_t(index))
}

// SetVertexBytes sets a block of data for the vertex function.
//
// Reference: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515846-setvertexbytes.
func (rce RenderCommandEncoder) SetVertexBytes(bytes unsafe.Pointer, length uintptr, index int) {
	C.RenderCommandEncoder_SetVertexBytes(rce.commandEncoder, bytes, C.size_t(length), C.uint_t(index))
}

// DrawPrimitives renders one instance of primitives using vertex data
// in contiguous array elements.
//
// Reference: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1516326-drawprimitives.
func (rce RenderCommandEncoder) DrawPrimitives(typ PrimitiveType, vertexStart, vertexCount int) {
	C.RenderCommandEncoder_DrawPrimitives(rce.commandEncoder, C.uint8_t(typ), C.uint_t(vertexStart), C.uint_t(vertexCount))
}

// BlitCommandEncoder is an encoder that specifies resource copy
// and resource synchronization commands.
//
// Reference: https://developer.apple.com/documentation/metal/mtlblitcommandencoder.
type BlitCommandEncoder struct {
	CommandEncoder
}

// CopyFromTexture encodes a command to copy image data from a slice of
// a source texture into a slice of a destination texture.
//
// Reference: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/1400754-copyfromtexture.
func (bce BlitCommandEncoder) CopyFromTexture(
	src Texture, srcSlice, srcLevel int, srcOrigin Origin, srcSize Size,
	dst Texture, dstSlice, dstLevel int, dstOrigin Origin,
) {
	C.BlitCommandEncoder_CopyFromTexture(
		bce.commandEncoder,
		src.texture, C.uint_t(srcSlice), C.uint_t(srcLevel),
		C.struct_Origin{
			X: C.uint_t(srcOrigin.X),
			Y: C.uint_t(srcOrigin.Y),
			Z: C.uint_t(srcOrigin.Z),
		},
		C.struct_Size{
			Width:  C.uint_t(srcSize.Width),
			Height: C.uint_t(srcSize.Height),
			Depth:  C.uint_t(srcSize.Depth),
		},
		dst.texture, C.uint_t(dstSlice), C.uint_t(dstLevel),
		C.struct_Origin{
			X: C.uint_t(dstOrigin.X),
			Y: C.uint_t(dstOrigin.Y),
			Z: C.uint_t(dstOrigin.Z),
		},
	)
}

// SynchronizeResource flushes any copy of the specified resource from its corresponding
// Device caches and, if needed, invalidates any CPU caches.
//
// Reference: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/1400775-synchronize.
func (bce BlitCommandEncoder) SynchronizeResource(resource Resource) {
	C.BlitCommandEncoder_SynchronizeResource(bce.commandEncoder, resource.resource())
}
