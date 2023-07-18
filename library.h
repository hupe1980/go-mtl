// +build darwin

#include "mtl.h"

struct Library {
	void *       Library;
	const char * Error;
};

struct Library Device_NewLibraryWithSource(void * device, const char * source, size_t sourceLength);

void * Library_NewFunctionWithName(void * library, const char * name);