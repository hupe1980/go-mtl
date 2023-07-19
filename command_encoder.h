// +build darwin

#include "mtl.h"

void CommandEncoder_EndEncoding(void * commandEncoder);

void ComputeCommandEncoder_SetComputePipelineState(void * computeCommandEncoder, void * computePipelineState);
void ComputeCommandEncoder_SetBuffer(void * computeCommandEncoder, void * buffer, uint_t offset, uint_t index);
void ComputeCommandEncoder_DispatchThreads(void * computeCommandEncoder, struct Size gridSize, struct Size threadgroupSize);

void RenderCommandEncoder_SetRenderPipelineState(void * renderCommandEncoder, void * renderPipelineState);
void RenderCommandEncoder_SetVertexBuffer(void * renderCommandEncoder, void * buffer, uint_t offset, uint_t index);
void RenderCommandEncoder_SetVertexBytes(void * renderCommandEncoder, const void * bytes, size_t length, uint_t index);
void RenderCommandEncoder_DrawPrimitives(void * renderCommandEncoder, uint8_t primitiveType, uint_t vertexStart, uint_t vertexCount);

void BlitCommandEncoder_CopyFromTexture(void * blitCommandEncoder,
	void * srcTexture, uint_t srcSlice, uint_t srcLevel, struct Origin srcOrigin, struct Size srcSize,
	void * dstTexture, uint_t dstSlice, uint_t dstLevel, struct Origin dstOrigin);
void BlitCommandEncoder_SynchronizeResource(void * blitCommandEncoder, void * resource);