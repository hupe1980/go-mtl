// +build darwin

#import <Metal/Metal.h>
#include "device.h"

struct Device CreateSystemDefaultDevice() {
	id<MTLDevice> device = MTLCreateSystemDefaultDevice();
	if (!device) {
		struct Device d;
		d.Device = NULL;
		return d;
	}

	struct Device d;
	d.Device = device;
	d.Headless = device.headless;
	d.LowPower = device.lowPower;
	d.Removable = device.removable;
	d.RegistryID = device.registryID;
	d.Name = device.name.UTF8String;
	
    return d;
}

// Caller must call free(d.devices).
struct Devices CopyAllDevices() {
	NSArray<id<MTLDevice>> * devices = MTLCopyAllDevices();

	struct Devices d;
	d.Devices = malloc(devices.count * sizeof(struct Device));
	for (int i = 0; i < devices.count; i++) {
		d.Devices[i].Device = devices[i];
		d.Devices[i].Headless = devices[i].headless;
		d.Devices[i].LowPower = devices[i].lowPower;
		d.Devices[i].Removable = devices[i].removable;
		d.Devices[i].RegistryID = devices[i].registryID;
		d.Devices[i].Name = devices[i].name.UTF8String;
	}
	
    d.Length = devices.count;
	
    return d;
}

bool Device_SupportsFamily(void * device, uint16_t gpuFamily) {
    return [(id<MTLDevice>)device supportsFamily:gpuFamily];
}

void * Device_NewBufferWithLength(void * device, size_t length, uint16_t options) {
	return [(id<MTLDevice>)device newBufferWithLength:(NSUInteger)length
	                                         options:(MTLResourceOptions)options];
}

void * Device_NewBufferWithBytes(void * device, const void * bytes, size_t length, uint16_t options) {
	return [(id<MTLDevice>)device newBufferWithBytes:(const void *)bytes
	                                          length:(NSUInteger)length
	                                         options:(MTLResourceOptions)options];
}

void * Buffer_Contents(void * buffer) {
    return ((id<MTLBuffer>)buffer).contents;
}