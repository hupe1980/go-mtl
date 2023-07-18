# 🍏 go-mtl: Go Bindings for Apple Metal
![Build Status](https://github.com/hupe1980/go-mtl/workflows/build/badge.svg) 
[![Go Reference](https://pkg.go.dev/badge/github.com/hupe1980/go-mtl.svg)](https://pkg.go.dev/github.com/hupe1980/go-mtl)
> go-mtl provides seamless integration between Go and Apple Metal, enabling developers to harness the full potential of Metal's high-performance graphics and compute capabilities in their Go applications. With go-mtl, you can write efficient and scalable code for darwin platforms, leveraging Metal's advanced features such as parallel processing, GPU acceleration, and low-level access to the graphics pipeline.

## Installation
Use Go modules to include go-mtl in your project:
```bash
go get github.com/hupe1980/go-mtl
```

## Usage
```go
import (
	"fmt"
	"log"
	"unsafe"

	"github.com/hupe1980/go-mtl"
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

func main() {
	device, err := mtl.CreateSystemDefaultDevice()
	if err != nil {
		log.Fatal(err)
	}

	lib, err := device.NewLibraryWithSource(source, mtl.CompileOptions{})
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

	b1 := device.NewBufferWithBytes(unsafe.Pointer(&dataA[0]), unsafe.Sizeof(dataA), mtl.ResourceStorageModeShared)
	b2 := device.NewBufferWithBytes(unsafe.Pointer(&dataB[0]), unsafe.Sizeof(dataB), mtl.ResourceStorageModeShared)
	r := device.NewBufferWithLength(unsafe.Sizeof(arrLen), mtl.ResourceStorageModeShared)

	cb := q.CommandBuffer()

	cce := cb.ComputeCommandEncoder()
	cce.SetComputePipelineState(pipelineState)
	cce.SetBuffer(b1, 0, 0)
	cce.SetBuffer(b2, 0, 1)
	cce.SetBuffer(r, 0, 2)

	cce.DispatchThreads(mtl.Size{Width: arrLen, Height: 1, Depth: 1}, mtl.Size{Width: 4, Height: 1, Depth: 1})

	cce.EndEncoding()

	cb.Commit()
	cb.WaitUntilCompleted()

	fmt.Println((*[1 << 30]float32)(r.Contents())[:arrLen])
}
```
Output:
```text
[0 2 4 6]
```

For more example usage, see [examples](./examples).

## Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request for any improvements or new features you would like to see.

## License
This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.