//go:build darwin
// +build darwin

package mtl

/*
#include "work_submission.h"
*/
import "C"
import (
	"unsafe"
)

// CommandQueue represents a queue that organizes the order
// in which command buffers are executed by the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandqueue
type CommandQueue struct {
	commandQueue unsafe.Pointer
}

// NewCommandQueue creates a new command queue for the specified device.
// The command queue is used to submit rendering and computation commands to the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/1433388-newcommandqueue
func (d Device) NewCommandQueue() CommandQueue {
	return CommandQueue{C.Device_NewCommandQueue(d.device)}
}

// CommandBuffer is a container that stores encoded commands
// that are committed to and executed by the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandbuffer
type CommandBuffer struct {
	commandBuffer unsafe.Pointer
}

// CommandBuffer returns a command buffer from the command queue that maintains strong references to resources.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandqueue/1508686-commandbuffer
func (cq CommandQueue) CommandBuffer() CommandBuffer {
	return CommandBuffer{C.CommandQueue_CommandBuffer(cq.commandQueue)}
}

// Commit submits the command buffer to run on the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443003-commit
func (cb CommandBuffer) Commit() {
	C.CommandBuffer_Commit(cb.commandBuffer)
}

// WaitUntilCompleted waits for the execution of this command buffer to complete.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443039-waituntilcompleted
func (cb CommandBuffer) WaitUntilCompleted() {
	C.CommandBuffer_WaitUntilCompleted(cb.commandBuffer)
}

func (cb CommandBuffer) ComputeCommandEncoder() ComputeCommandEncoder {
	return ComputeCommandEncoder{CommandEncoder{C.CommandBuffer_ComputeCommandEncoder(cb.commandBuffer)}}
}

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
