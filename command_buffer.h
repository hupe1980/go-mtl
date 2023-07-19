// +build darwin

#include "mtl.h"

struct RenderPassDescriptor {
	uint8_t           ColorAttachment0LoadAction;
	uint8_t           ColorAttachment0StoreAction;
	struct ClearColor ColorAttachment0ClearColor;
	void *            ColorAttachment0Texture;
};

void   CommandBuffer_Commit(void * commandBuffer);
void   CommandBuffer_WaitUntilCompleted(void * commandBuffer);
void   CommandBuffer_PresentDrawable(void * commandBuffer, void * drawable);
void * CommandBuffer_ComputeCommandEncoder(void * commandBuffer);
void * CommandBuffer_RenderCommandEncoderWithDescriptor(void * commandBuffer, struct RenderPassDescriptor descriptor);
void * CommandBuffer_BlitCommandEncoder(void * commandBuffer);