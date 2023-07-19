// +build darwin

#import <Metal/Metal.h>
#include "library.h"

struct Library Device_NewLibraryWithSource(void * device, const char * source, size_t sourceLength, struct CompileOptions opts) {
	MTLCompileOptions *compileOptions = [MTLCompileOptions new];
	compileOptions.fastMathEnabled = opts.FastMathEnabled;
	compileOptions.preserveInvariance = opts.PreserveInvariance;

	NSError * error;
	id<MTLLibrary> library = [(id<MTLDevice>)device
		newLibraryWithSource:[[NSString alloc] initWithBytes:source length:sourceLength encoding:NSUTF8StringEncoding]
		options:compileOptions
		error:&error];

	struct Library l;
	l.Library = library;
	if (!library) {
        l.Error = error.localizedDescription.UTF8String;
	}

	return l;
}

void * Library_NewFunctionWithName(void * library, const char * name) {
	return [(id<MTLLibrary>)library newFunctionWithName:[NSString stringWithUTF8String:name]];
}
