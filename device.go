//go:build darwin
// +build darwin

package mtl

/*
#include <stdlib.h>
#include <stdbool.h>
#include "device.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

// Device is an abstract representation of the GPU and serves as the primary interface for a Metal app.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice
type Device struct {
	device unsafe.Pointer

	// Headless indicates whether a device is configured as headless.
	Headless bool

	// LowPower indicates whether a device is low-power.
	LowPower bool

	// Removable determines whether or not a GPU is removable.
	Removable bool

	// RegistryID is the registry ID value for the device.
	RegistryID uint64

	// Name is the name of the device.
	Name string
}

// CreateSystemDefaultDevice returns the preferred system default Metal device.
//
// Reference: https://developer.apple.com/documentation/metal/1433401-mtlcreatesystemdefaultdevice
func CreateSystemDefaultDevice() (Device, error) {
	d := C.CreateSystemDefaultDevice()
	if d.Device == nil {
		return Device{}, errors.New("metal is not supported on this system")
	}

	return Device{
		device:     d.Device,
		Headless:   bool(d.Headless),
		LowPower:   bool(d.LowPower),
		Removable:  bool(d.Removable),
		RegistryID: uint64(d.RegistryID),
		Name:       C.GoString(d.Name),
	}, nil
}

// CopyAllDevices returns all Metal devices in the system.
//
// Reference: https://developer.apple.com/documentation/metal/1433367-mtlcopyalldevices
func CopyAllDevices() []Device {
	d := C.CopyAllDevices()
	defer C.free(unsafe.Pointer(d.Devices))

	ds := make([]Device, d.Length)
	for i := 0; i < len(ds); i++ {
		d := (*C.struct_Device)(unsafe.Pointer(uintptr(unsafe.Pointer(d.Devices)) + uintptr(i)*C.sizeof_struct_Device))

		ds[i].device = d.Device
		ds[i].Headless = bool(d.Headless)
		ds[i].LowPower = bool(d.LowPower)
		ds[i].Removable = bool(d.Removable)
		ds[i].RegistryID = uint64(d.RegistryID)
		ds[i].Name = C.GoString(d.Name)
	}

	return ds
}

// Device returns the underlying id<MTLDevice> pointer.
func (d Device) Device() unsafe.Pointer {
	return d.device
}

// SupportsFamily reports whether the device supports the feature set of the GPU family.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/3143473-supportsfamily
func (d Device) SupportsFamily(gf GPUFamily) bool {
	return bool(C.Device_SupportsFamily(d.device, C.uint16_t(gf)))
}
