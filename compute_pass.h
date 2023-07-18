// +build darwin

struct ComputePipelineState {
	void * 			ComputePipelineState;
	const char * 	Error;
};

struct ComputePipelineState Device_NewComputePipelineStateWithFunction(void * device, void * function);