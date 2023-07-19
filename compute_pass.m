// +build darwin

#import <Metal/Metal.h>
#include "compute_pass.h"

struct ComputePipelineState Device_NewComputePipelineStateWithFunction(void * device,  void * function) {
    NSError * error; 
    id<MTLComputePipelineState> pipelineState = [(id<MTLDevice>)device newComputePipelineStateWithFunction:function error:&error];

    struct ComputePipelineState cps;
	cps.ComputePipelineState = pipelineState;
	
	if (!pipelineState) {
		cps.Error = error.localizedDescription.UTF8String;
		
		return cps;
	}

	cps.MaxTotalThreadsPerThreadgroup = pipelineState.maxTotalThreadsPerThreadgroup;

	return cps;
}