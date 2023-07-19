// +build darwin

#import <Metal/Metal.h>
#include "command_encoder.h"


void CommandEncoder_EndEncoding(void * commandEncoder) {
	[(id<MTLCommandEncoder>)commandEncoder endEncoding];
}

void ComputeCommandEncoder_SetComputePipelineState(void * computeCommandEncoder, void * computePipelineState) {
	[(id<MTLComputeCommandEncoder>)computeCommandEncoder setComputePipelineState:(id<MTLComputePipelineState>)computePipelineState];
}

void ComputeCommandEncoder_SetBuffer(void * computeCommandEncoder, void * buffer, uint_t offset, uint_t index) {
	[(id<MTLComputeCommandEncoder>)computeCommandEncoder setBuffer:(id<MTLBuffer>)buffer
	                                                            offset:(NSUInteger)offset
	                                                           atIndex:(NSUInteger)index];
}

void ComputeCommandEncoder_DispatchThreads(void * computeCommandEncoder, struct Size gridSize, struct Size threadgroupSize) {
    [(id<MTLComputeCommandEncoder>)computeCommandEncoder dispatchThreads:(MTLSize){gridSize.Width, gridSize.Height, gridSize.Depth}
                                                                    threadsPerThreadgroup:(MTLSize){threadgroupSize.Width, threadgroupSize.Height, threadgroupSize.Depth}];   
}

void RenderCommandEncoder_SetRenderPipelineState(void * renderCommandEncoder, void * renderPipelineState) {
	[(id<MTLRenderCommandEncoder>)renderCommandEncoder setRenderPipelineState:(id<MTLRenderPipelineState>)renderPipelineState];
}

void RenderCommandEncoder_SetVertexBuffer(void * renderCommandEncoder, void * buffer, uint_t offset, uint_t index) {
	[(id<MTLRenderCommandEncoder>)renderCommandEncoder setVertexBuffer:(id<MTLBuffer>)buffer
	                                                            offset:(NSUInteger)offset
	                                                           atIndex:(NSUInteger)index];
}

void RenderCommandEncoder_SetVertexBytes(void * renderCommandEncoder, const void * bytes, size_t length, uint_t index) {
	[(id<MTLRenderCommandEncoder>)renderCommandEncoder setVertexBytes:bytes
	                                                           length:(NSUInteger)length
	                                                          atIndex:(NSUInteger)index];
}

void RenderCommandEncoder_DrawPrimitives(void * renderCommandEncoder, uint8_t primitiveType, uint_t vertexStart, uint_t vertexCount) {
	[(id<MTLRenderCommandEncoder>)renderCommandEncoder drawPrimitives:(MTLPrimitiveType)primitiveType
	                                                      vertexStart:(NSUInteger)vertexStart
	                                                      vertexCount:(NSUInteger)vertexCount];
}

void BlitCommandEncoder_CopyFromTexture(void * blitCommandEncoder,
	void * srcTexture, uint_t srcSlice, uint_t srcLevel, struct Origin srcOrigin, struct Size srcSize,
	void * dstTexture, uint_t dstSlice, uint_t dstLevel, struct Origin dstOrigin) {
	[(id<MTLBlitCommandEncoder>)blitCommandEncoder copyFromTexture:(id<MTLTexture>)srcTexture
	                                                   sourceSlice:(NSUInteger)srcSlice
	                                                   sourceLevel:(NSUInteger)srcLevel
	                                                  sourceOrigin:(MTLOrigin){srcOrigin.X, srcOrigin.Y, srcOrigin.Z}
	                                                    sourceSize:(MTLSize){srcSize.Width, srcSize.Height, srcSize.Depth}
	                                                     toTexture:(id<MTLTexture>)dstTexture
	                                              destinationSlice:(NSUInteger)dstSlice
	                                              destinationLevel:(NSUInteger)dstLevel
	                                             destinationOrigin:(MTLOrigin){dstOrigin.X, dstOrigin.Y, dstOrigin.Z}];
}

void BlitCommandEncoder_SynchronizeResource(void * blitCommandEncoder, void * resource) {
	[(id<MTLBlitCommandEncoder>)blitCommandEncoder synchronizeResource:(id<MTLResource>)resource];
}
