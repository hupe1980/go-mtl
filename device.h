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
