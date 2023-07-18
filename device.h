// +build darwin

#include "mtl.h"

struct Device {
	void *       Device;
	bool         Headless;
	bool         LowPower;
	bool         Removable;
	uint64_t     RegistryID;
	const char * Name;
};

struct Devices {
	struct Device * Devices;
	int             Length;
};


struct Device CreateSystemDefaultDevice();
struct Devices CopyAllDevices();

bool   Device_SupportsFamily(void * device, uint16_t gpuFamily);
void * Device_NewBufferWithLength(void * device, size_t length, uint16_t options);
void * Device_NewBufferWithBytes(void * device, const void * bytes, size_t length, uint16_t options);

void * Buffer_Contents(void * buffer);
