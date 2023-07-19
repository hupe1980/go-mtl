// +build darwin

#include "mtl.h"

struct Library {
	void *       Library;
	const char * Error;
};

struct CompileOptions {
	bool FastMathEnabled;
	bool PreserveInvariance;
	uint_t LanguageVersion;
};

struct Library Device_NewLibraryWithSource(void * device, const char * source, size_t sourceLength, struct CompileOptions opts);

void * Library_NewFunctionWithName(void * library, const char * name);