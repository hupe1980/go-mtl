// +build darwin

#import <Metal/Metal.h>
#include "command_buffer.h"

void CommandBuffer_Commit(void * commandBuffer) {
	[(id<MTLCommandBuffer>)commandBuffer commit];
}

void CommandBuffer_WaitUntilCompleted(void * commandBuffer) {
	[(id<MTLCommandBuffer>)commandBuffer waitUntilCompleted];
}

void CommandBuffer_PresentDrawable(void * commandBuffer, void * drawable) {
	[(id<MTLCommandBuffer>)commandBuffer presentDrawable:(id<MTLDrawable>)drawable];
}

void * CommandBuffer_ComputeCommandEncoder(void * commandBuffer) {
	return [(id<MTLCommandBuffer>)commandBuffer computeCommandEncoder];
}

void * CommandBuffer_RenderCommandEncoderWithDescriptor(void * commandBuffer, struct RenderPassDescriptor descriptor) {
	MTLRenderPassDescriptor * renderPassDescriptor = [[MTLRenderPassDescriptor alloc] init];
	renderPassDescriptor.colorAttachments[0].loadAction = descriptor.ColorAttachment0LoadAction;
	renderPassDescriptor.colorAttachments[0].storeAction = descriptor.ColorAttachment0StoreAction;
	renderPassDescriptor.colorAttachments[0].clearColor = MTLClearColorMake(
		descriptor.ColorAttachment0ClearColor.Red,
		descriptor.ColorAttachment0ClearColor.Green,
		descriptor.ColorAttachment0ClearColor.Blue,
		descriptor.ColorAttachment0ClearColor.Alpha
	);
	renderPassDescriptor.colorAttachments[0].texture = (id<MTLTexture>)descriptor.ColorAttachment0Texture;
	return [(id<MTLCommandBuffer>)commandBuffer renderCommandEncoderWithDescriptor:renderPassDescriptor];
}

void * CommandBuffer_BlitCommandEncoder(void * commandBuffer) {
	return [(id<MTLCommandBuffer>)commandBuffer blitCommandEncoder];
}
