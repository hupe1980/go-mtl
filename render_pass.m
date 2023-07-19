// +build darwin

#import <Metal/Metal.h>
#include "render_pass.h"

struct RenderPipelineState Device_NewRenderPipelineStateWithDescriptor(void * device, struct RenderPipelineDescriptor descriptor) {
	MTLRenderPipelineDescriptor * renderPipelineDescriptor = [[MTLRenderPipelineDescriptor alloc] init];
	renderPipelineDescriptor.vertexFunction = descriptor.VertexFunction;
	renderPipelineDescriptor.fragmentFunction = descriptor.FragmentFunction;
	renderPipelineDescriptor.colorAttachments[0].pixelFormat = descriptor.ColorAttachment0PixelFormat;

	NSError * error;
	id<MTLRenderPipelineState> renderPipelineState = [(id<MTLDevice>)device newRenderPipelineStateWithDescriptor:renderPipelineDescriptor
	                                                                                                       error:&error];
                                                                                                           
	struct RenderPipelineState rps;
	rps.RenderPipelineState = renderPipelineState;
	if (!renderPipelineState) {
		rps.Error = error.localizedDescription.UTF8String;
        return rps;
	}

	return rps;
}