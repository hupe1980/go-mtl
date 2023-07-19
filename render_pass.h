// +build darwin

#include "mtl.h"

struct RenderPipelineDescriptor {
	void *   VertexFunction;
	void *   FragmentFunction;
	uint16_t ColorAttachment0PixelFormat;
};

struct RenderPipelineState {
	void *       RenderPipelineState;
	const char * Error;
};

struct RenderPipelineState 	Device_NewRenderPipelineStateWithDescriptor(void * device, struct RenderPipelineDescriptor descriptor);