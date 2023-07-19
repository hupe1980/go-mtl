// +build darwin

#import <Metal/Metal.h>
#include "command_queue.h"

void * Device_NewCommandQueue(void * device) {
	return [(id<MTLDevice>)device newCommandQueue];
}

void * CommandQueue_CommandBuffer(void * commandQueue) {
	return [(id<MTLCommandQueue>)commandQueue commandBuffer];
}
