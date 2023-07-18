//go:build darwin
// +build darwin

package mtl

/*
#include "library.h"
struct Library Go_Device_NewLibraryWithSource(void * device, _GoString_ source) {
	return Device_NewLibraryWithSource(device, _GoStringPtr(source), _GoStringLen(source));
}
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

// CompileOptions specifies optional compilation settings for
// the graphics or compute functions within a library.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcompileoptions
type CompileOptions struct {
	// TODO.
}

// Library represents a collection of compiled graphics or compute functions.
//
// Reference: https://developer.apple.com/documentation/metal/mtllibrary.
type Library struct {
	library unsafe.Pointer
}

// NewLibraryWithSource creates a new library that contains
// the functions stored in the specified source string.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/1433431-newlibrarywithsource
func (d Device) NewLibraryWithSource(source string, opt CompileOptions) (Library, error) {
	l := C.Go_Device_NewLibraryWithSource(d.device, source) // TODO: opt.
	if l.Library == nil {
		return Library{}, errors.New(C.GoString(l.Error))
	}

	return Library{l.Library}, nil
}

// Function represents a programmable graphics or compute function executed by the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtlfunction.
type Function struct {
	function unsafe.Pointer
}

// NewFunctionWithName creates a new function object that represents a shader function in the library.
//
// Reference: https://developer.apple.com/documentation/metal/mtllibrary/1515524-newfunctionwithname
func (l Library) NewFunctionWithName(name string) (Function, error) {
	f := C.Library_NewFunctionWithName(l.library, C.CString(name))
	if f == nil {
		return Function{}, fmt.Errorf("function %q not found", name)
	}

	return Function{f}, nil
}
