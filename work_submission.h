// +build darwin

#include "mtl.h"

void * Device_NewCommandQueue(void * device);

void * CommandQueue_CommandBuffer(void * commandQueue);

void   CommandBuffer_Commit(void * commandBuffer);
void   CommandBuffer_WaitUntilCompleted(void * commandBuffer);
void * CommandBuffer_ComputeCommandEncoder(void * commandBuffer);

void CommandEncoder_EndEncoding(void * commandEncoder);

void ComputeCommandEncoder_SetComputePipelineState(void * computeCommandEncoder, void * computePipelineState);
void ComputeCommandEncoder_SetBuffer(void * computeCommandEncoder, void * buffer, uint_t offset, uint_t index);
void ComputeCommandEncoder_DispatchThreads(void * computeCommandEncoder, struct Size gridSize, struct Size threadgroupSize);