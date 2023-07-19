//go:build darwin
// +build darwin

package mtl

/*
#include "compute_pass.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

// ComputePipelineState represents an object that contains a compiled compute pipeline.
//
// Referece: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate
type ComputePipelineState struct {
	computePipelineState          unsafe.Pointer
	MaxTotalThreadsPerThreadgroup uint
}

// NewComputePipelineStateWithFunction creates a new ComputePipelineState object with the specified compute function.
// It takes a Function object representing the compute function to be used and returns a ComputePipelineState object.
// An error is returned if the creation fails.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/1433395-newcomputepipelinestatewithfunct
func (d Device) NewComputePipelineStateWithFunction(f Function) (ComputePipelineState, error) {
	cps := C.Device_NewComputePipelineStateWithFunction(d.device, f.function)
	if cps.ComputePipelineState == nil {
		return ComputePipelineState{}, errors.New(C.GoString(cps.Error))
	}

	return ComputePipelineState{
		computePipelineState:          cps.ComputePipelineState,
		MaxTotalThreadsPerThreadgroup: uint(cps.MaxTotalThreadsPerThreadgroup),
	}, nil
}
