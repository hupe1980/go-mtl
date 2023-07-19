package mtl

import (
	"os"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

const source = `#include <metal_stdlib>

using namespace metal;

kernel void add_arrays(device const float* inA,
	device const float* inB,
	device float* result,
	uint index [[thread_position_in_grid]])
{
// the for-loop is replaced with a collection of threads, each of which
// calls this function.
result[index] = inA[index] + inB[index];
}
`

func TestCalculation(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		// GPU functions are not available for macOS runners
		// https://github.com/actions/runner-images/issues/1779#issuecomment-707071183
		t.Skip()
	}

	// Create a Metal device.
	device, err := CreateSystemDefaultDevice()
	assert.NoError(t, err)

	// Create a Metal library from the provided source code.
	lib, err := device.NewLibraryWithSource(source, CompileOptions{})
	assert.NoError(t, err)

	// Retrieve the Metal function named "add_arrays" from the library.
	addArrays, err := lib.NewFunctionWithName("add_arrays")
	assert.NoError(t, err)

	// Create a Metal compute pipeline state with the function.
	pipelineState, err := device.NewComputePipelineStateWithFunction(addArrays)
	assert.NoError(t, err)

	// Create a Metal command queue to submit commands for execution.
	q := device.NewCommandQueue()

	// Set the length of the arrays.
	arrLen := uint(4)

	// Prepare the input data.
	dataA := []float32{0.0, 1.0, 2.0, 3.0}
	dataB := []float32{0.0, 1.0, 2.0, 3.0}

	// Create Metal buffers for input and output data.
	// b1 and b2 represent the input arrays, and r represents the output array.
	b1 := device.NewBufferWithBytes(unsafe.Pointer(&dataA[0]), unsafe.Sizeof(dataA), ResourceStorageModeShared)
	b2 := device.NewBufferWithBytes(unsafe.Pointer(&dataB[0]), unsafe.Sizeof(dataB), ResourceStorageModeShared)
	r := device.NewBufferWithLength(unsafe.Sizeof(arrLen), ResourceStorageModeShared)

	// // Create a Metal command buffer to encode and execute commands.
	cb := q.CommandBuffer()

	// Create a compute command encoder to encode compute commands.
	cce := cb.ComputeCommandEncoder()

	// Set the compute pipeline state to specify the function to be executed.
	cce.SetComputePipelineState(pipelineState)

	// Set the input and output buffers for the compute function.
	cce.SetBuffer(b1, 0, 0)
	cce.SetBuffer(b2, 0, 1)
	cce.SetBuffer(r, 0, 2)

	// Dispatch compute threads to perform the calculation.
	cce.DispatchThreads(Size{Width: arrLen, Height: 1, Depth: 1}, Size{Width: 4, Height: 1, Depth: 1})

	// End encoding the compute command.
	cce.EndEncoding()

	// Commit the command buffer for execution.
	cb.Commit()

	// Wait until the command buffer execution is completed.
	cb.WaitUntilCompleted()

	// Read the results from the output buffer
	result := (*[1 << 30]float32)(r.Contents())[:arrLen]

	assert.ElementsMatch(t, []float32{0.0, 2.0, 4.0, 6.0}, result)
}
