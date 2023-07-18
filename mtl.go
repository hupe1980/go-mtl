//go:build darwin
// +build darwin

// Package mtl provides Go bindings for the Metal framework.
package mtl

/*
#cgo LDFLAGS: -framework Metal -framework CoreGraphics -framework Foundation
#include "mtl.h"
*/
import "C"

// GPUFamily is a family of GPUs.
//
// Reference: https://developer.apple.com/documentation/metal/mtlgpufamily.
type GPUFamily uint16

const (
	GPUFamilyApple1 GPUFamily = 1001 // Apple family 1 GPU features that correspond to the Apple A7 GPUs.
	GPUFamilyApple2 GPUFamily = 1002 // Apple family 2 GPU features that correspond to the Apple A8 GPUs.
	GPUFamilyApple3 GPUFamily = 1003 // Apple family 3 GPU features that correspond to the Apple A9 and A10 GPUs.
	GPUFamilyApple4 GPUFamily = 1004 // Apple family 4 GPU features that correspond to the Apple A11 GPUs.
	GPUFamilyApple5 GPUFamily = 1005 // Apple family 5 GPU features that correspond to the Apple A12 GPUs.
	GPUFamilyApple6 GPUFamily = 1006 // Apple family 6 GPU features that correspond to the Apple A13 GPUs.
	GPUFamilyApple7 GPUFamily = 1007 // Apple family 7 GPU features that correspond to the Apple A14 and M1 GPUs.
	GPUFamilyApple8 GPUFamily = 1008 // Apple family 8 GPU features that correspond to the Apple A15 and M2 GPUs.
	//GPUFamilyMac1    GPUFamily = 2001 // Mac family 1 GPU features.
	GPUFamilyMac2    GPUFamily = 2002 // Mac family 2 GPU features.
	GPUFamilyCommon1 GPUFamily = 3001 // Common family 1 GPU features.
	GPUFamilyCommon2 GPUFamily = 3002 // Common family 2 GPU features.
	GPUFamilyCommon3 GPUFamily = 3003 // Common family 3 GPU features.
	GPUFamilyMetal3  GPUFamily = 5001 // Metal 3 features.
)

// CPUCacheMode is the CPU cache mode that defines the CPU mapping of a resource.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcpucachemode.
type CPUCacheMode uint8

const (
	// CPUCacheModeDefaultCache is the default CPU cache mode for the resource.
	// Guarantees that read and write operations are executed in the expected order.
	CPUCacheModeDefaultCache CPUCacheMode = 0

	// CPUCacheModeWriteCombined is a write-combined CPU cache mode for the resource.
	// Optimized for resources that the CPU will write into, but never read.
	CPUCacheModeWriteCombined CPUCacheMode = 1
)

// StorageMode defines defines the memory location and access permissions of a resource.
//
// Reference: https://developer.apple.com/documentation/metal/mtlstoragemode.
type StorageMode uint8

const (
	// StorageModeShared indicates that the resource is stored in system memory
	// accessible to both the CPU and the GPU.
	StorageModeShared StorageMode = 0

	// StorageModeManaged indicates that the resource exists as a synchronized
	// memory pair with one copy stored in system memory accessible to the CPU
	// and another copy stored in video memory accessible to the GPU.
	StorageModeManaged StorageMode = 1

	// StorageModePrivate indicates that the resource is stored in memory
	// only accessible to the GPU. In iOS and tvOS, the resource is stored in
	// system memory. In macOS, the resource is stored in video memory.
	StorageModePrivate StorageMode = 2

	// StorageModeMemoryless indicates that the resource is stored in on-tile memory,
	// without CPU or GPU memory backing. The contents of the on-tile memory are undefined
	// and do not persist; the only way to populate the resource is to render into it.
	// Memoryless resources are limited to temporary render targets (i.e., Textures configured
	// with a TextureDescriptor and used with a RenderPassAttachmentDescriptor).
	StorageModeMemoryless StorageMode = 3
)

const (
	resourceCPUCacheModeShift       = 0
	resourceStorageModeShift        = 4
	resourceHazardTrackingModeShift = 8
)

// ResourceOptions defines optional arguments used to create
// and influence behavior of buffer and texture objects.
//
// Reference: https://developer.apple.com/documentation/metal/mtlresourceoptions.
type ResourceOptions uint16

const (
	// ResourceCPUCacheModeDefaultCache is the default CPU cache mode for the resource.
	// Guarantees that read and write operations are executed in the expected order.
	ResourceCPUCacheModeDefaultCache ResourceOptions = ResourceOptions(CPUCacheModeDefaultCache) << resourceCPUCacheModeShift

	// ResourceCPUCacheModeWriteCombined is a write-combined CPU cache mode for the resource.
	// Optimized for resources that the CPU will write into, but never read.
	ResourceCPUCacheModeWriteCombined ResourceOptions = ResourceOptions(CPUCacheModeWriteCombined) << resourceCPUCacheModeShift

	// ResourceStorageModeShared indicates that the resource is stored in system memory
	// accessible to both the CPU and the GPU.
	ResourceStorageModeShared ResourceOptions = ResourceOptions(StorageModeShared) << resourceStorageModeShift

	// ResourceStorageModeManaged indicates that the resource exists as a synchronized
	// memory pair with one copy stored in system memory accessible to the CPU
	// and another copy stored in video memory accessible to the GPU.
	ResourceStorageModeManaged ResourceOptions = ResourceOptions(StorageModeManaged) << resourceStorageModeShift

	// ResourceStorageModePrivate indicates that the resource is stored in memory
	// only accessible to the GPU. In iOS and tvOS, the resource is stored
	// in system memory. In macOS, the resource is stored in video memory.
	ResourceStorageModePrivate ResourceOptions = ResourceOptions(StorageModePrivate) << resourceStorageModeShift

	// ResourceStorageModeMemoryless indicates that the resource is stored in on-tile memory,
	// without CPU or GPU memory backing. The contents of the on-tile memory are undefined
	// and do not persist; the only way to populate the resource is to render into it.
	// Memoryless resources are limited to temporary render targets (i.e., Textures configured
	// with a TextureDescriptor and used with a RenderPassAttachmentDescriptor).
	ResourceStorageModeMemoryless ResourceOptions = ResourceOptions(StorageModeMemoryless) << resourceStorageModeShift

	// ResourceHazardTrackingModeUntracked indicates that the command encoder dependencies
	// for this resource are tracked manually with Fence objects. This value is always set
	// for resources sub-allocated from a Heap object and may optionally be specified for
	// non-heap resources.
	ResourceHazardTrackingModeUntracked ResourceOptions = 1 << resourceHazardTrackingModeShift
)

// Size represents the set of dimensions that declare the size of an object,
// such as an image, texture, threadgroup, or grid.
//
// Reference: https://developer.apple.com/documentation/metal/mtlsize.
type Size struct{ Width, Height, Depth uint }
