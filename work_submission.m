// +build darwin

#import <Metal/Metal.h>
#include "work_submission.h"

void * Device_NewCommandQueue(void * device) {
	return [(id<MTLDevice>)device newCommandQueue];
}

void * CommandQueue_CommandBuffer(void * commandQueue) {
	return [(id<MTLCommandQueue>)commandQueue commandBuffer];
}

void CommandBuffer_Commit(void * commandBuffer) {
	[(id<MTLCommandBuffer>)commandBuffer commit];
}

void CommandBuffer_WaitUntilCompleted(void * commandBuffer) {
	[(id<MTLCommandBuffer>)commandBuffer waitUntilCompleted];
}

void * CommandBuffer_ComputeCommandEncoder(void * commandBuffer) {
	return [(id<MTLCommandBuffer>)commandBuffer computeCommandEncoder];
}

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