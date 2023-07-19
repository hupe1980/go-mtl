//go:build darwin
// +build darwin

package mtl

/*
#include "render_pass.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

// RenderPipelineColorAttachmentDescriptor represents a color attachment descriptor for a render pipeline.
type RenderPipelineColorAttachmentDescriptor struct {
	// PixelFormat is the pixel format of the color attachment's texture.
	PixelFormat PixelFormat
}

// RenderPipelineDescriptor represents a descriptor for a render pipeline.
type RenderPipelineDescriptor struct {
	// VertexFunction is a programmable function that processes individual vertices in a rendering pass.
	VertexFunction Function

	// FragmentFunction is a programmable function that processes individual fragments in a rendering pass.
	FragmentFunction Function

	// ColorAttachments is an array of attachments that store color data.
	ColorAttachments [1]RenderPipelineColorAttachmentDescriptor
}

// RenderPipelineState represents the state of a render pipeline.
type RenderPipelineState struct {
	renderPipelineState unsafe.Pointer
}

// NewRenderPipelineStateWithDescriptor creates a new render pipeline state with the provided descriptor.
// It returns the created RenderPipelineState or an error if the creation fails.
func (d Device) NewRenderPipelineStateWithDescriptor(rpd RenderPipelineDescriptor) (RenderPipelineState, error) {
	descriptor := C.struct_RenderPipelineDescriptor{
		VertexFunction:              rpd.VertexFunction.function,
		FragmentFunction:            rpd.FragmentFunction.function,
		ColorAttachment0PixelFormat: C.uint16_t(rpd.ColorAttachments[0].PixelFormat),
	}

	rps := C.Device_NewRenderPipelineStateWithDescriptor(d.device, descriptor)

	if rps.RenderPipelineState == nil {
		return RenderPipelineState{}, errors.New(C.GoString(rps.Error))
	}

	return RenderPipelineState{rps.RenderPipelineState}, nil
}
