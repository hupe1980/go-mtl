package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"unsafe"

	"github.com/hupe1980/go-mtl"
)

// The source code for the Metal shaders.
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

func main() {
	// Create a Metal device.
	device, err := mtl.CreateSystemDefaultDevice()
	if err != nil {
		log.Fatal(err)
	}

	// Create a Metal library from the provided source code.
	lib, err := device.NewLibraryWithSource(source)
	if err != nil {
		log.Fatal(err)
	}

	// Get the vertex shader function from the library.
	vs, err := lib.NewFunctionWithName("vertex_shader")
	if err != nil {
		log.Fatal(err)
	}

	// Get the fragment shader function from the library.
	fs, err := lib.NewFunctionWithName("fragment_shader")
	if err != nil {
		log.Fatal(err)
	}

	// Create a render pipeline state descriptor and configure it.
	var rpld mtl.RenderPipelineDescriptor
	rpld.VertexFunction = vs
	rpld.FragmentFunction = fs
	rpld.ColorAttachments[0].PixelFormat = mtl.PixelFormatBGRA8Unorm

	// Create a render pipeline state object from the descriptor.
	rps, err := device.NewRenderPipelineStateWithDescriptor(rpld)
	if err != nil {
		log.Fatal(err)
	}

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
	vertexBuffer := device.NewBufferWithBytes(unsafe.Pointer(&vertexData[0]), unsafe.Sizeof(vertexData), mtl.ResourceStorageModeManaged)

	// Create an output texture to render into.
	td := mtl.TextureDescriptor{
		PixelFormat: mtl.PixelFormatBGRA8Unorm,
		Width:       256,
		Height:      256,
		StorageMode: mtl.StorageModeManaged,
	}

	texture := device.NewTextureWithDescriptor(td)

	// Create a command queue for the device.
	cq := device.NewCommandQueue()

	// Create a command buffer.
	cb := cq.CommandBuffer()

	// Encode all render commands.
	var rpd mtl.RenderPassDescriptor
	rpd.ColorAttachments[0].LoadAction = mtl.LoadActionClear
	rpd.ColorAttachments[0].StoreAction = mtl.StoreActionStore
	rpd.ColorAttachments[0].ClearColor = mtl.ClearColor{Red: 0.35, Green: 0.65, Blue: 0.85, Alpha: 1}
	rpd.ColorAttachments[0].Texture = texture

	// Create a render command encoder with the render pass descriptor.
	rce := cb.RenderCommandEncoderWithDescriptor(rpd)

	// Set the render pipeline state object.
	rce.SetRenderPipelineState(rps)

	// Set the vertex buffer.
	rce.SetVertexBuffer(vertexBuffer, 0, 0)

	// Draw the triangle.
	rce.DrawPrimitives(mtl.PrimitiveTypeTriangle, 0, 3)

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
	region := mtl.RegionMake2D(0, 0, texture.Width, texture.Height)
	texture.GetBytes(&img.Pix[0], uintptr(bytesPerRow), region, 0)

	// Save the image to a PNG file.
	file, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generated image saved to output.png")
}
