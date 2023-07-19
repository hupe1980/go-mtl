package mtl

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestCalculation(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		// GPU functions are not available for macOS runners
		// https://github.com/actions/runner-images/issues/1779#issuecomment-707071183
		t.Skip()
	}

	// Create a Metal device.
	device, err := CreateSystemDefaultDevice()
	require.NoError(t, err)

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

	// Create a Metal library from the provided source code.
	lib, err := device.NewLibraryWithSource(source)
	require.NoError(t, err)

	// Retrieve the Metal function named "add_arrays" from the library.
	addArrays, err := lib.NewFunctionWithName("add_arrays")
	require.NoError(t, err)

	// Create a Metal compute pipeline state with the function.
	pipelineState, err := device.NewComputePipelineStateWithFunction(addArrays)
	require.NoError(t, err)

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

	// Specify threadgroup size
	tgs := pipelineState.MaxTotalThreadsPerThreadgroup
	if tgs > arrLen {
		tgs = arrLen
	}

	// Dispatch compute threads to perform the calculation.
	cce.DispatchThreads(Size{Width: arrLen, Height: 1, Depth: 1}, Size{Width: tgs, Height: 1, Depth: 1})

	// End encoding the compute command.
	cce.EndEncoding()

	// Commit the command buffer for execution.
	cb.Commit()

	// Wait until the command buffer execution is completed.
	cb.WaitUntilCompleted()

	// Read the results from the output buffer
	result := (*[1 << 30]float32)(r.Contents())[:arrLen]

	require.ElementsMatch(t, []float32{0.0, 2.0, 4.0, 6.0}, result)
}

func TestRender(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		// GPU functions are not available for macOS runners
		// https://github.com/actions/runner-images/issues/1779#issuecomment-707071183
		t.Skip()
	}

	// Create a Metal device.
	device, err := CreateSystemDefaultDevice()
	require.NoError(t, err)

	const source = `#include <metal_stdlib>

using namespace metal;

// Define the vertex structure with position and color attributes.
struct Vertex {
	float4 position [[position]];
	float4 color;
};

// Vertex shader function that returns the vertex data.
vertex Vertex vertex_shader(
	uint vertexID [[vertex_id]],
	device Vertex * vertices [[buffer(0)]]
) {
	return vertices[vertexID];
}

// Fragment shader function that returns the color for each fragment.
fragment float4 fragment_shader(Vertex in [[stage_in]]) {
	return in.color;
}
`

	// Create a Metal library from the provided source code.
	lib, err := device.NewLibraryWithSource(source)
	require.NoError(t, err)

	// Get the vertex shader function from the library.
	vs, err := lib.NewFunctionWithName("vertex_shader")
	require.NoError(t, err)

	// Get the fragment shader function from the library.
	fs, err := lib.NewFunctionWithName("fragment_shader")
	require.NoError(t, err)

	// Create a render pipeline state descriptor and configure it.
	var rpld RenderPipelineDescriptor
	rpld.VertexFunction = vs
	rpld.FragmentFunction = fs
	rpld.ColorAttachments[0].PixelFormat = PixelFormatBGRA8Unorm

	// Create a render pipeline state object from the descriptor.
	rps, err := device.NewRenderPipelineStateWithDescriptor(rpld)
	require.NoError(t, err)

	// Define the vertex data for the triangle.
	type Vertex struct {
		Position [4]float32
		Color    [4]float32
	}

	vertexData := [...]Vertex{
		{Position: [4]float32{+0.00, +0.5, 0, 1}, Color: [4]float32{1, 0, 0, 1}},
		{Position: [4]float32{-0.5, -0.5, 0, 1}, Color: [4]float32{0, 1, 0, 1}},
		{Position: [4]float32{+0.5, -0.5, 0, 1}, Color: [4]float32{0, 0, 1, 1}},
	}

	// Create a vertex buffer with the vertex data.
	vertexBuffer := device.NewBufferWithBytes(unsafe.Pointer(&vertexData[0]), unsafe.Sizeof(vertexData), ResourceStorageModeManaged)

	// Create an output texture to render into.
	td := TextureDescriptor{
		PixelFormat: PixelFormatBGRA8Unorm,
		Width:       256,
		Height:      256,
		StorageMode: StorageModeManaged,
	}

	texture := device.NewTextureWithDescriptor(td)

	// Create a command queue for the device.
	cq := device.NewCommandQueue()

	// Create a command buffer.
	cb := cq.CommandBuffer()

	// Encode all render commands.
	var rpd RenderPassDescriptor
	rpd.ColorAttachments[0].LoadAction = LoadActionClear
	rpd.ColorAttachments[0].StoreAction = StoreActionStore
	rpd.ColorAttachments[0].ClearColor = ClearColor{Red: 0.35, Green: 0.65, Blue: 0.85, Alpha: 1}
	rpd.ColorAttachments[0].Texture = texture

	// Create a render command encoder with the render pass descriptor.
	rce := cb.RenderCommandEncoderWithDescriptor(rpd)

	// Set the render pipeline state object.
	rce.SetRenderPipelineState(rps)

	// Set the vertex buffer.
	rce.SetVertexBuffer(vertexBuffer, 0, 0)

	// Draw the triangle.
	rce.DrawPrimitives(PrimitiveTypeTriangle, 0, 3)

	// End encoding of the render commands.
	rce.EndEncoding()

	// Encode all blit commands.
	bce := cb.BlitCommandEncoder()

	// Synchronize the output texture.
	bce.SynchronizeResource(texture)

	// End encoding of the blit commands.
	bce.EndEncoding()

	// Commit the command buffer.
	cb.Commit()

	// Wait until the command buffer is completed.
	cb.WaitUntilCompleted()

	// Read pixels from output texture into an image.
	img := image.NewNRGBA(image.Rect(0, 0, int(texture.Width), int(texture.Height)))
	bytesPerRow := 4 * texture.Width
	region := RegionMake2D(0, 0, texture.Width, texture.Height)
	texture.GetBytes(&img.Pix[0], uintptr(bytesPerRow), region, 0)

	// Open file to compare
	want, err := readPNG("testdata/triangle.png")
	require.NoError(t, err)

	require.Equal(t, img.Bounds(), want.Bounds())

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c, _ := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
			d, _ := color.NRGBAModel.Convert(want.At(x, y)).(color.NRGBA)

			require.Equal(t, c.R, d.R)
			require.Equal(t, c.G, d.G)
			require.Equal(t, c.B, d.B)
			require.Equal(t, c.A, d.A)
		}
	}
}

func readPNG(name string) (image.Image, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return png.Decode(f)
}
