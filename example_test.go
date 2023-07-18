package mtl

import (
	"fmt"
	"log"
	"unsafe"
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

func Example_calculation() {
	device, err := CreateSystemDefaultDevice()
	if err != nil {
		log.Fatal(err)
	}

	lib, err := device.NewLibraryWithSource(source, CompileOptions{})
	if err != nil {
		log.Fatal(err)
	}

	addArrays, err := lib.NewFunctionWithName("add_arrays")
	if err != nil {
		log.Fatal(err)
	}

	pipelineState, err := device.NewComputePipelineStateWithFunction(addArrays)
	if err != nil {
		log.Fatal(err)
	}

	q := device.NewCommandQueue()

	arrLen := uint(4)

	dataA := []float32{0.0, 1.0, 2.0, 3.0}
	dataB := []float32{0.0, 1.0, 2.0, 3.0}

	b1 := device.NewBufferWithBytes(unsafe.Pointer(&dataA[0]), unsafe.Sizeof(dataA), ResourceStorageModeShared)
	b2 := device.NewBufferWithBytes(unsafe.Pointer(&dataB[0]), unsafe.Sizeof(dataB), ResourceStorageModeShared)
	r := device.NewBufferWithLength(unsafe.Sizeof(arrLen), ResourceStorageModeShared)

	cb := q.CommandBuffer()

	cce := cb.ComputeCommandEncoder()
	cce.SetComputePipelineState(pipelineState)
	cce.SetBuffer(b1, 0, 0)
	cce.SetBuffer(b2, 0, 1)
	cce.SetBuffer(r, 0, 2)

	cce.DispatchThreads(Size{Width: arrLen, Height: 1, Depth: 1}, Size{Width: 4, Height: 1, Depth: 1})

	cce.EndEncoding()

	cb.Commit()
	cb.WaitUntilCompleted()

	fmt.Println((*[1 << 30]float32)(r.Contents())[:arrLen])

	// Output: [0 2 4 6]
}
