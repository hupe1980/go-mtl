//go:build darwin
// +build darwin

// Package mtl provides Go bindings for the Metal framework.
package mtl

/*
#cgo LDFLAGS: -framework Metal -framework CoreGraphics -framework Foundation
#include "mtl.h"
*/
import "C"
import "unsafe"

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

// StorageMode represents the memory location and access permissions for a resource.
//
// Reference: https://developer.apple.com/documentation/metal/mtlstoragemode
type StorageMode uint8

const (
	// StorageModeShared indicates that the resource is stored in system memory and is accessible to both the CPU and the GPU.
	StorageModeShared StorageMode = 0

	// StorageModeManaged indicates that the CPU and GPU may maintain separate copies of the resource,
	// and any changes must be explicitly synchronized.
	StorageModeManaged StorageMode = 1

	// StorageModePrivate indicates that the resource can be accessed only by the GPU.
	StorageModePrivate StorageMode = 2

	// StorageModeMemoryless indicates that the resource's contents can be accessed only by the GPU
	// and only exist temporarily during a render pass.
	StorageModeMemoryless StorageMode = 3
)

// HazardTrackingMode represents the options for hazard tracking mode in Metal.
type HazardTrackingMode uint8

// Hazard tracking modes.
const (
	// HazardTrackingModeDefault specifies the default hazard tracking mode.
	HazardTrackingModeDefault HazardTrackingMode = 0

	// HazardTrackingModeUntracked specifies that hazards must be prevented manually when modifying object contents.
	HazardTrackingModeUntracked HazardTrackingMode = 1

	// HazardTrackingModeTracked specifies that Metal automatically prevents hazards when modifying object contents.
	HazardTrackingModeTracked HazardTrackingMode = 2
)

const (
	resourceCPUCacheModeShift       = 0
	resourceStorageModeShift        = 4
	resourceHazardTrackingModeShift = 8
)

// ResourceOptions defines optional arguments used to set the behavior of a resource.
//
// Reference: https://developer.apple.com/documentation/metal/mtlresourceoptions
type ResourceOptions uint16

const (
	// ResourceCPUCacheModeDefaultCache specifies the default CPU cache mode for the resource.
	// Guarantees that read and write operations are executed in the expected order.
	ResourceCPUCacheModeDefaultCache ResourceOptions = ResourceOptions(CPUCacheModeDefaultCache) << resourceCPUCacheModeShift

	// ResourceCPUCacheModeWriteCombined specifies a write-combined CPU cache mode that is optimized for resources
	// that the CPU writes into, but never reads.
	ResourceCPUCacheModeWriteCombined ResourceOptions = ResourceOptions(CPUCacheModeWriteCombined) << resourceCPUCacheModeShift

	// ResourceStorageModeShared indicates that the resource is stored in system memory
	// and is accessible to both the CPU and the GPU.
	ResourceStorageModeShared ResourceOptions = ResourceOptions(StorageModeShared) << resourceStorageModeShift

	// ResourceStorageModeManaged indicates that the CPU and GPU may maintain separate copies of the resource,
	// which need to be explicitly synchronized.
	ResourceStorageModeManaged ResourceOptions = ResourceOptions(StorageModeManaged) << resourceStorageModeShift

	// ResourceStorageModePrivate indicates that the resource can be accessed only by the GPU.
	ResourceStorageModePrivate ResourceOptions = ResourceOptions(StorageModePrivate) << resourceStorageModeShift

	// ResourceStorageModeMemoryless indicates that the resource's contents can be accessed only by the GPU
	// and only exist temporarily during a render pass.
	ResourceStorageModeMemoryless ResourceOptions = ResourceOptions(StorageModeMemoryless) << resourceStorageModeShift

	// ResourceHazardTrackingModeDefault specifies that the default hazard tracking mode should be used.
	ResourceHazardTrackingModeDefault ResourceOptions = ResourceOptions(HazardTrackingModeDefault) << resourceHazardTrackingModeShift

	// ResourceHazardTrackingModeTracked specifies that Metal prevents hazards when modifying this object's contents.
	ResourceHazardTrackingModeTracked ResourceOptions = ResourceOptions(HazardTrackingModeTracked) << resourceHazardTrackingModeShift

	// ResourceHazardTrackingModeUntracked specifies that the app must prevent hazards when modifying this object's contents.
	ResourceHazardTrackingModeUntracked ResourceOptions = ResourceOptions(HazardTrackingModeUntracked) << resourceHazardTrackingModeShift
)

// LanguageVersion represents different versions of the Metal shading language.
type LanguageVersion uint32

const (
	// LanguageVersion1_1 represents version 1.1 of the Metal shading language.
	LanguageVersion1_1 LanguageVersion = (1 << 16) + 1

	// LanguageVersion1_2 represents version 1.2 of the Metal shading language.
	LanguageVersion1_2 LanguageVersion = (1 << 16) + 2

	// LanguageVersion2_0 represents version 2.0 of the Metal shading language.
	LanguageVersion2_0 LanguageVersion = (2 << 16)

	// LanguageVersion2_1 represents version 2.1 of the Metal shading language.
	LanguageVersion2_1 LanguageVersion = (2 << 16) + 1

	// LanguageVersion2_2 represents version 2.1 of the Metal shading language.
	LanguageVersion2_2 LanguageVersion = (2 << 16) + 2

	// LanguageVersion2_3 represents version 2.1 of the Metal shading language.
	LanguageVersion2_3 LanguageVersion = (2 << 16) + 3

	// LanguageVersion2_4 represents version 2.1 of the Metal shading language.
	LanguageVersion2_4 LanguageVersion = (2 << 16) + 4

	// LanguageVersion3_0 represents version 2.1 of the Metal shading language.
	LanguageVersion3_0 LanguageVersion = (3 << 16) + 0
)

// PixelFormat represents the data formats that describe the organization and characteristics of individual pixels in a texture.
type PixelFormat uint16

const (
	// Ordinary 8-Bit Pixel Formats

	// PixelFormatA8Unorm is an ordinary format with one 8-bit normalized unsigned integer component.
	PixelFormatA8Unorm PixelFormat = 1

	// PixelFormatR8Unorm is an ordinary format with one 8-bit normalized unsigned integer component.
	PixelFormatR8Unorm PixelFormat = 10

	// PixelFormatR8UnormSRGB is an ordinary format with one 8-bit normalized unsigned integer component with conversion between sRGB and linear space.
	PixelFormatR8UnormSRGB PixelFormat = 11

	// PixelFormatR8Snorm is an ordinary format with one 8-bit normalized signed integer component.
	PixelFormatR8Snorm PixelFormat = 12

	// PixelFormatR8Uint is an ordinary format with one 8-bit unsigned integer component.
	PixelFormatR8Uint PixelFormat = 13

	// PixelFormatR8Sint is an ordinary format with one 8-bit signed integer component.
	PixelFormatR8Sint PixelFormat = 14

	// Ordinary 16-Bit Pixel Formats

	// PixelFormatR16Unorm is an ordinary format with one 16-bit normalized unsigned integer component.
	PixelFormatR16Unorm PixelFormat = 20

	// PixelFormatR16Snorm is an ordinary format with one 16-bit normalized signed integer component.
	PixelFormatR16Snorm PixelFormat = 22

	// PixelFormatR16Uint is an ordinary format with one 16-bit unsigned integer component.
	PixelFormatR16Uint PixelFormat = 23

	// PixelFormatR16Sint is an ordinary format with one 16-bit signed integer component.
	PixelFormatR16Sint PixelFormat = 24

	// PixelFormatR16Float is an ordinary format with one 16-bit floating-point component.
	PixelFormatR16Float PixelFormat = 25

	// PixelFormatRG8Unorm is an ordinary format with two 8-bit normalized unsigned integer components.
	PixelFormatRG8Unorm PixelFormat = 30

	// PixelFormatRG8UnormSRGB is an ordinary format with two 8-bit normalized unsigned integer components with conversion between sRGB and linear space.
	PixelFormatRG8UnormSRGB PixelFormat = 31

	// PixelFormatRG8Snorm is an ordinary format with two 8-bit normalized signed integer components.
	PixelFormatRG8Snorm PixelFormat = 32

	// PixelFormatRG8Uint is an ordinary format with two 8-bit unsigned integer components.
	PixelFormatRG8Uint PixelFormat = 33

	// PixelFormatRG8Sint is an ordinary format with two 8-bit signed integer components.
	PixelFormatRG8Sint PixelFormat = 34

	// Packed 16-Bit Pixel Formats

	// PixelFormatB5G6R5Unorm is a packed 16-bit format with normalized unsigned integer color components: 5 bits for blue, 6 bits for green, 5 bits for red, packed into 16 bits.
	PixelFormatB5G6R5Unorm PixelFormat = 40

	// PixelFormatA1BGR5Unorm is a packed 16-bit format with normalized unsigned integer color components: 5 bits each for BGR and 1 for alpha, packed into 16 bits.
	PixelFormatA1BGR5Unorm PixelFormat = 41

	// PixelFormatABGR4Unorm is a packed 16-bit format with normalized unsigned integer color components: 4 bits each for ABGR, packed into 16 bits.
	PixelFormatABGR4Unorm PixelFormat = 42

	// PixelFormatBGR5A1Unorm is a packed 16-bit format with normalized unsigned integer color components: 5 bits each for BGR and 1 for alpha, packed into 16 bits.
	PixelFormatBGR5A1Unorm PixelFormat = 43

	// Ordinary 32-Bit Pixel Formats

	// PixelFormatR32Uint is an ordinary format with one 32-bit unsigned integer component.
	PixelFormatR32Uint PixelFormat = 53

	// PixelFormatR32Sint is an ordinary format with one 32-bit signed integer component.
	PixelFormatR32Sint PixelFormat = 54

	// PixelFormatR32Float is an ordinary format with one 32-bit floating-point component.
	PixelFormatR32Float PixelFormat = 55

	// PixelFormatRG16Unorm is an ordinary format with two 16-bit normalized unsigned integer components.
	PixelFormatRG16Unorm PixelFormat = 60

	// PixelFormatRG16Snorm is an ordinary format with two 16-bit normalized signed integer components.
	PixelFormatRG16Snorm PixelFormat = 62

	// PixelFormatRG16Uint is an ordinary format with two 16-bit unsigned integer components.
	PixelFormatRG16Uint PixelFormat = 63

	// PixelFormatRG16Sint is an ordinary format with two 16-bit signed integer components.
	PixelFormatRG16Sint PixelFormat = 64

	// PixelFormatRG16Float is an ordinary format with two 16-bit floating-point components.
	PixelFormatRG16Float PixelFormat = 65

	// PixelFormatRGBA8Unorm is an ordinary format with four 8-bit normalized unsigned integer components in RGBA order.
	PixelFormatRGBA8Unorm PixelFormat = 70

	// PixelFormatRGBA8UnormSRGB is an ordinary format with four 8-bit normalized unsigned integer components in RGBA order with conversion between sRGB and linear space.
	PixelFormatRGBA8UnormSRGB PixelFormat = 71

	// PixelFormatRGBA8Snorm is an ordinary format with four 8-bit normalized signed integer components in RGBA order.
	PixelFormatRGBA8Snorm PixelFormat = 72

	// PixelFormatRGBA8Uint is an ordinary format with four 8-bit unsigned integer components in RGBA order.
	PixelFormatRGBA8Uint PixelFormat = 73

	// PixelFormatRGBA8Sint is an ordinary format with four 8-bit signed integer components in RGBA order.
	PixelFormatRGBA8Sint PixelFormat = 74

	// PixelFormatBGRA8Unorm is an ordinary format with four 8-bit normalized unsigned integer components in BGRA order.
	PixelFormatBGRA8Unorm PixelFormat = 80

	// PixelFormatBGRA8UnormSRGB is an ordinary format with four 8-bit normalized unsigned integer components in BGRA order with conversion between sRGB and linear space.
	PixelFormatBGRA8UnormSRGB PixelFormat = 81

	// Packed 32-Bit Pixel Formats

	// PixelFormatBGR10A2Unorm is a packed 32-bit format with normalized unsigned integer color components: 10 bits for blue, green, and red, and 2 bits for alpha, packed into 32 bits.
	PixelFormatBGR10A2Unorm PixelFormat = 94

	// PixelFormatRGB10A2Unorm is a packed 32-bit format with normalized unsigned integer color components: 10 bits for red, green, and blue, and 2 bits for alpha, packed into 32 bits.
	PixelFormatRGB10A2Unorm PixelFormat = 90

	// PixelFormatRGB10A2Uint is a packed 32-bit format with unsigned integer color components: 10 bits for red, green, and blue, and 2 bits for alpha, packed into 32 bits.
	PixelFormatRGB10A2Uint PixelFormat = 91

	// PixelFormatRG11B10Float is a packed 32-bit format with floating-point color components: 11 bits for red and green, and 10 bits for blue, packed into 32 bits.
	PixelFormatRG11B10Float PixelFormat = 92

	// PixelFormatRGB9E5Float is a packed 32-bit format with floating-point color components: 9 bits for red, green, and blue, and 5 bits for exponent, packed into 32 bits.
	PixelFormatRGB9E5Float PixelFormat = 93

	// Ordinary 64-Bit Pixel Formats

	// PixelFormatRG32Uint is an ordinary format with two 32-bit unsigned integer components.
	PixelFormatRG32Uint PixelFormat = 103

	// PixelFormatRG32Sint is an ordinary format with two 32-bit signed integer components.
	PixelFormatRG32Sint PixelFormat = 104

	// PixelFormatRG32Float is an ordinary format with two 32-bit floating-point components.
	PixelFormatRG32Float PixelFormat = 105

	// PixelFormatRGBA16Unorm is an ordinary format with four 16-bit normalized unsigned integer components in RGBA order.
	PixelFormatRGBA16Unorm PixelFormat = 110

	// PixelFormatRGBA16Snorm is an ordinary format with four 16-bit normalized signed integer components in RGBA order.
	PixelFormatRGBA16Snorm PixelFormat = 112

	// PixelFormatRGBA16Uint is an ordinary format with four 16-bit unsigned integer components in RGBA order.
	PixelFormatRGBA16Uint PixelFormat = 113

	// PixelFormatRGBA16Sint is an ordinary format with four 16-bit signed integer components in RGBA order.
	PixelFormatRGBA16Sint PixelFormat = 114

	// PixelFormatRGBA16Float is an ordinary format with four 16-bit floating-point components in RGBA order.
	PixelFormatRGBA16Float PixelFormat = 115

	// Ordinary 128-Bit Pixel Formats

	// PixelFormatRGBA32Uint is an ordinary format with four 32-bit unsigned integer components in RGBA order.
	PixelFormatRGBA32Uint PixelFormat = 123

	// PixelFormatRGBA32Sint is an ordinary format with four 32-bit signed integer components in RGBA order.
	PixelFormatRGBA32Sint PixelFormat = 124

	// PixelFormatRGBA32Float is an ordinary format with four 32-bit floating-point components in RGBA order.
	PixelFormatRGBA32Float PixelFormat = 125

	// Compressed PVRTC Pixel Formats

	// PixelFormatPVRTCRGB2BPP is a compressed PVRTC format with RGB colors and 2 bits per pixel.
	PixelFormatPVRTCRGB2BPP PixelFormat = 160

	// PixelFormatPVRTCRGB2BPPSRGB is a compressed PVRTC format with RGB colors and 2 bits per pixel, with conversion between sRGB and linear space.
	PixelFormatPVRTCRGB2BPPSRGB PixelFormat = 161

	// PixelFormatPVRTCRGB4BPP is a compressed PVRTC format with RGB colors and 4 bits per pixel.
	PixelFormatPVRTCRGB4BPP PixelFormat = 162

	// PixelFormatPVRTCRGB4BPPSRGB is a compressed PVRTC format with RGB colors and 4 bits per pixel, with conversion between sRGB and linear space.
	PixelFormatPVRTCRGB4BPPSRGB PixelFormat = 163

	// PixelFormatPVRTCRGBA2BPP is a compressed PVRTC format with RGBA colors and 2 bits per pixel.
	PixelFormatPVRTCRGBA2BPP PixelFormat = 164

	// PixelFormatPVRTCRGBA2BPPSRGB is a compressed PVRTC format with RGBA colors and 2 bits per pixel, with conversion between sRGB and linear space.
	PixelFormatPVRTCRGBA2BPPSRGB PixelFormat = 165

	// PixelFormatPVRTCRGBA4BPP is a compressed PVRTC format with RGBA colors and 4 bits per pixel.
	PixelFormatPVRTCRGBA4BPP PixelFormat = 166

	// PixelFormatPVRTCRGBA4BPPSRGB is a compressed PVRTC format with RGBA colors and 4 bits per pixel, with conversion between sRGB and linear space.
	PixelFormatPVRTCRGBA4BPPSRGB PixelFormat = 167

	// Compressed EAC/ETC Pixel Formats

	// PixelFormatEACR11Unorm is a compressed EAC format with one 11-bit normalized unsigned integer component.
	PixelFormatEACR11Unorm PixelFormat = 170

	// PixelFormatEACR11Snorm is a compressed EAC format with one 11-bit normalized signed integer component.
	PixelFormatEACR11Snorm PixelFormat = 172

	// PixelFormatEACRG11Unorm is a compressed EAC format with two 11-bit normalized unsigned integer components.
	PixelFormatEACRG11Unorm PixelFormat = 174

	// PixelFormatEACRG11Snorm is a compressed EAC format with two 11-bit normalized signed integer components.
	PixelFormatEACRG11Snorm PixelFormat = 176

	// PixelFormatEACRGBA8 is a compressed EAC format with RGBA colors and 8 bits per pixel.
	PixelFormatEACRGBA8 PixelFormat = 178

	// PixelFormatEACRGBA8SRGB is a compressed EAC format with RGBA colors and 8 bits per pixel, with conversion between sRGB and linear space.
	PixelFormatEACRGBA8SRGB PixelFormat = 179

	// PixelFormatETC2RGB8 is a compressed ETC2 format with RGB colors and 8 bits per pixel.
	PixelFormatETC2RGB8 PixelFormat = 180

	// PixelFormatETC2RGB8SRGB is a compressed ETC2 format with RGB colors and 8 bits per pixel, with conversion between sRGB and linear space.
	PixelFormatETC2RGB8SRGB PixelFormat = 181

	// PixelFormatETC2RGB8A1 is a compressed ETC2 format with RGB colors and 8 bits per pixel, with 1-bit alpha.
	PixelFormatETC2RGB8A1 PixelFormat = 182

	// PixelFormatETC2RGB8A1SRGB is a compressed ETC2 format with RGB colors and 8 bits per pixel, with 1-bit alpha, and conversion between sRGB and linear space.
	PixelFormatETC2RGB8A1SRGB PixelFormat = 183

	// Compressed ASTC Pixel Formats

	// PixelFormatASTC4x4SRGB is a compressed ASTC format with 4x4 blocks and sRGB color space.
	PixelFormatASTC4x4SRGB PixelFormat = 186

	// PixelFormatASTC5x4SRGB is a compressed ASTC format with 5x4 blocks and sRGB color space.
	PixelFormatASTC5x4SRGB PixelFormat = 187

	// PixelFormatASTC5x5SRGB is a compressed ASTC format with 5x5 blocks and sRGB color space.
	PixelFormatASTC5x5SRGB PixelFormat = 188

	// PixelFormatASTC6x5SRGB is a compressed ASTC format with 6x5 blocks and sRGB color space.
	PixelFormatASTC6x5SRGB PixelFormat = 189

	// PixelFormatASTC6x6SRGB is a compressed ASTC format with 6x6 blocks and sRGB color space.
	PixelFormatASTC6x6SRGB PixelFormat = 190

	// PixelFormatASTC8x5SRGB is a compressed ASTC format with 8x5 blocks and sRGB color space.
	PixelFormatASTC8x5SRGB PixelFormat = 192

	// PixelFormatASTC8x6SRGB is a compressed ASTC format with 8x6 blocks and sRGB color space.
	PixelFormatASTC8x6SRGB PixelFormat = 193

	// PixelFormatASTC8x8SRGB is a compressed ASTC format with 8x8 blocks and sRGB color space.
	PixelFormatASTC8x8SRGB PixelFormat = 194

	// PixelFormatASTC10x5SRGB is a compressed ASTC format with 10x5 blocks and sRGB color space.
	PixelFormatASTC10x5SRGB PixelFormat = 195

	// PixelFormatASTC10x6SRGB is a compressed ASTC format with 10x6 blocks and sRGB color space.
	PixelFormatASTC10x6SRGB PixelFormat = 196

	// PixelFormatASTC10x8SRGB is a compressed ASTC format with 10x8 blocks and sRGB color space.
	PixelFormatASTC10x8SRGB PixelFormat = 197

	// PixelFormatASTC10x10SRGB is a compressed ASTC format with 10x10 blocks and sRGB color space.
	PixelFormatASTC10x10SRGB PixelFormat = 198

	// PixelFormatASTC12x10SRGB is a compressed ASTC format with 12x10 blocks and sRGB color space.
	PixelFormatASTC12x10SRGB PixelFormat = 199

	// PixelFormatASTC12x12SRGB is a compressed ASTC format with 12x12 blocks and sRGB color space.
	PixelFormatASTC12x12SRGB PixelFormat = 200

	// PixelFormatASTC4x4LDR is an ASTC pixel format with low-dynamic-range content, a block width of 4, and a block height of 4.
	PixelFormatASTC4x4LDR PixelFormat = 204

	// PixelFormatASTC5x4LDR is an ASTC pixel format with low-dynamic-range content, a block width of 5, and a block height of 4.
	PixelFormatASTC5x4LDR PixelFormat = 205

	// PixelFormatASTC5x5LDR is an ASTC pixel format with low-dynamic-range content, a block width of 5, and a block height of 5.
	PixelFormatASTC5x5LDR PixelFormat = 206

	// PixelFormatASTC6x5LDR is an ASTC pixel format with low-dynamic-range content, a block width of 6, and a block height of 5.
	PixelFormatASTC6x5LDR PixelFormat = 207

	// PixelFormatASTC6x6LDR is an ASTC pixel format with low-dynamic-range content, a block width of 6, and a block height of 6.
	PixelFormatASTC6x6LDR PixelFormat = 208

	// PixelFormatASTC8x5LDR is an ASTC pixel format with low-dynamic-range content, a block width of 8, and a block height of 5.
	PixelFormatASTC8x5LDR PixelFormat = 210

	// PixelFormatASTC8x6LDR is an ASTC pixel format with low-dynamic-range content, a block width of 8, and a block height of 6.
	PixelFormatASTC8x6LDR PixelFormat = 211

	// PixelFormatASTC8x8LDR is an ASTC pixel format with low-dynamic-range content, a block width of 8, and a block height of 8.
	PixelFormatASTC8x8LDR PixelFormat = 212

	// PixelFormatASTC10x5LDR is an ASTC pixel format with low-dynamic-range content, a block width of 10, and a block height of 5.
	PixelFormatASTC10x5LDR PixelFormat = 213

	// PixelFormatASTC10x6LDR is an ASTC pixel format with low-dynamic-range content, a block width of 10, and a block height of 6.
	PixelFormatASTC10x6LDR PixelFormat = 214

	// PixelFormatASTC10x8LDR is an ASTC pixel format with low-dynamic-range content, a block width of 10, and a block height of 8.
	PixelFormatASTC10x8LDR PixelFormat = 215

	// PixelFormatASTC10x10LDR is an ASTC pixel format with low-dynamic-range content,
	// a block width of 10, and a block height of 10.
	PixelFormatASTC10x10LDR PixelFormat = 216

	// PixelFormatASTC12x10LDR is an ASTC pixel format with low-dynamic-range content,
	// a block width of 12, and a block height of 10.
	PixelFormatASTC12x10LDR PixelFormat = 217

	// PixelFormatASTC12x12LDR is an ASTC pixel format with low-dynamic-range content,
	// a block width of 12, and a block height of 12.
	PixelFormatASTC12x12LDR PixelFormat = 218

	// PixelFormatASTC4x4HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 4, and a block height of 4.
	PixelFormatASTC4x4HDR PixelFormat = 222

	// PixelFormatASTC5x4HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 5, and a block height of 4.
	PixelFormatASTC5x4HDR PixelFormat = 223

	// PixelFormatASTC5x5HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 5, and a block height of 6.
	PixelFormatASTC5x5HDR PixelFormat = 224

	// PixelFormatASTC6x5HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 6, and a block height of 5.
	PixelFormatASTC6x5HDR PixelFormat = 225

	// PixelFormatASTC6x6HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 6, and a block height of 6.
	PixelFormatASTC6x6HDR PixelFormat = 226

	// PixelFormatASTC8x5HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 8, and a block height of 5.
	PixelFormatASTC8x5HDR PixelFormat = 228

	// PixelFormatASTC8x6HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 8, and a block height of 6.
	PixelFormatASTC8x6HDR PixelFormat = 229

	// PixelFormatASTC8x8HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 8, and a block height of 8.
	PixelFormatASTC8x8HDR PixelFormat = 230

	// PixelFormatASTC10x5HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 10, and a block height of 5.
	PixelFormatASTC10x5HDR PixelFormat = 231

	// PixelFormatASTC10x6HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 10, and a block height of 6.
	PixelFormatASTC10x6HDR PixelFormat = 232

	// PixelFormatASTC10x8HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 10, and a block height of 8.
	PixelFormatASTC10x8HDR PixelFormat = 233

	// PixelFormatASTC10x10HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 10, and a block height of 10.
	PixelFormatASTC10x10HDR PixelFormat = 234

	// PixelFormatASTC12x10HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 12, and a block height of 10.
	PixelFormatASTC12x10HDR PixelFormat = 235

	// PixelFormatASTC12x12HDR is an ASTC pixel format with high-dynamic-range content,
	// a block width of 12, and a block height of 12.
	PixelFormatASTC12x12HDR PixelFormat = 236

	// PixelFormatBC1RGBA is a compressed format with two 16-bit color components and one 32-bit descriptor component.
	PixelFormatBC1RGBA PixelFormat = 130

	// PixelFormatBC1RGBASRGB is a compressed format with two 16-bit color components and one 32-bit descriptor component,
	// with conversion between sRGB and linear space.
	PixelFormatBC1RGBASRGB PixelFormat = 131

	// PixelFormatBC2RGBA is a compressed format with two 64-bit chunks. The first chunk contains two 8-bit alpha components
	// and one 48-bit descriptor component. The second chunk contains two 16-bit color components and one 32-bit descriptor component.
	PixelFormatBC2RGBA PixelFormat = 132

	// PixelFormatBC2RGBASRGB is a compressed format with two 64-bit chunks, with conversion between sRGB and linear space.
	// The first chunk contains two 8-bit alpha components and one 48-bit descriptor component.
	// The second chunk contains two 16-bit color components and one 32-bit descriptor component.
	PixelFormatBC2RGBASRGB PixelFormat = 133

	// PixelFormatBC3RGBA is a compressed format with two 64-bit chunks. The first chunk contains two 8-bit alpha components
	// and one 48-bit descriptor component. The second chunk contains two 16-bit color components and one 32-bit descriptor component.
	PixelFormatBC3RGBA PixelFormat = 134

	// PixelFormatBC3RGBASRGB is a compressed format with two 64-bit chunks, with conversion between sRGB and linear space.
	// The first chunk contains two 8-bit alpha components and one 48-bit descriptor component.
	// The second chunk contains two 16-bit color components and one 32-bit descriptor component.
	PixelFormatBC3RGBASRGB PixelFormat = 135

	// PixelFormatBC4RUnorm is a compressed format with one normalized unsigned integer component.
	PixelFormatBC4RUnorm PixelFormat = 140

	// PixelFormatBC4RSnorm is a compressed format with one normalized signed integer component.
	PixelFormatBC4RSnorm PixelFormat = 141

	// PixelFormatBC5RGUnorm is a compressed format with two normalized unsigned integer components.
	PixelFormatBC5RGUnorm PixelFormat = 142

	// PixelFormatBC5RGSnorm is a compressed format with two normalized signed integer components.
	PixelFormatBC5RGSnorm PixelFormat = 143

	// PixelFormatBC6HRGBFloat is a compressed format with four floating-point components.
	PixelFormatBC6HRGBFloat PixelFormat = 150

	// PixelFormatBC6HRGBUfloat is a compressed format with four unsigned floating-point components.
	PixelFormatBC6HRGBUfloat PixelFormat = 151

	// PixelFormatBC7RGBAUnorm is a compressed format with four normalized unsigned integer components.
	PixelFormatBC7RGBAUnorm PixelFormat = 152

	// PixelFormatBC7RGBAUnormSRGB is a compressed format with four normalized unsigned integer components,
	// with conversion between sRGB and linear space.
	PixelFormatBC7RGBAUnormSRGB PixelFormat = 153

	// PixelFormatGBGR422 is a pixel format where the red and green components are subsampled horizontally.
	PixelFormatGBGR422 PixelFormat = 240

	// PixelFormatBGRG422 is a pixel format where the red and green components are subsampled horizontally.
	PixelFormatBGRG422 PixelFormat = 241

	// PixelFormatDepth16Unorm is a pixel format with a 16-bit normalized unsigned integer component, used for a depth render target.
	PixelFormatDepth16Unorm PixelFormat = 250

	// PixelFormatDepth32Float is a pixel format with one 32-bit floating-point component, used for a depth render target.
	PixelFormatDepth32Float PixelFormat = 252

	// PixelFormatStencil8 is a pixel format with an 8-bit unsigned integer component, used for a stencil render target.
	PixelFormatStencil8 PixelFormat = 253

	// PixelFormatDepth24UnormStencil8 is a 32-bit combined depth and stencil pixel format with a 24-bit normalized unsigned integer for depth and an 8-bit unsigned integer for stencil.
	PixelFormatDepth24UnormStencil8 PixelFormat = 255

	// PixelFormatDepth32FloatStencil8 is a 40-bit combined depth and stencil pixel format with a 32-bit floating-point value for depth and an 8-bit unsigned integer for stencil.
	PixelFormatDepth32FloatStencil8 PixelFormat = 260

	// PixelFormatX32Stencil8 is a stencil pixel format used to read the stencil value from a texture with a combined 32-bit depth and 8-bit stencil value.
	PixelFormatX32Stencil8 PixelFormat = 261

	// PixelFormatX24Stencil8 is a stencil pixel format used to read the stencil value from a texture with a combined 24-bit depth and 8-bit stencil value.
	PixelFormatX24Stencil8 PixelFormat = 262

	// PixelFormatBGRA10XR is a 64-bit extended range pixel format with four fixed-point components: 10-bit blue, 10-bit green, 10-bit red, and 10-bit alpha.
	PixelFormatBGRA10XR PixelFormat = 552

	// PixelFormatBGRA10XRSRGB is a 64-bit extended range pixel format with sRGB conversion and four fixed-point components: 10-bit blue, 10-bit green, 10-bit red, and 10-bit alpha.
	PixelFormatBGRA10XRSRGB PixelFormat = 553

	// PixelFormatBGR10XR is a 32-bit extended range pixel format with three fixed-point components: 10-bit blue, 10-bit green, and 10-bit red.
	PixelFormatBGR10XR PixelFormat = 554

	// PixelFormatBGR10XRSRGB is a 32-bit extended range pixel format with sRGB conversion and three fixed-point components: 10-bit blue, 10-bit green, and 10-bit red.
	PixelFormatBGR10XRSRGB PixelFormat = 555
)

// PrimitiveType defines geometric primitive types for drawing commands.
//
// Reference: https://developer.apple.com/documentation/metal/mtlprimitivetype
type PrimitiveType uint8

// Geometric primitive types for drawing commands.
const (
	// PrimitiveTypePoint rasterizes a point at each vertex. The vertex shader must provide [[point_size]], or the point size is undefined.
	PrimitiveTypePoint PrimitiveType = 0

	// PrimitiveTypeLine rasterizes a line between each separate pair of vertices, resulting in a series of unconnected lines. If there are an odd number of vertices, the last vertex is ignored.
	PrimitiveTypeLine PrimitiveType = 1

	// PrimitiveTypeLineStrip rasterizes a line between each pair of adjacent vertices, resulting in a series of connected lines (also called a polyline).
	PrimitiveTypeLineStrip PrimitiveType = 2

	// PrimitiveTypeTriangle rasterizes a triangle for every separate set of three vertices. If the number of vertices is not a multiple of three, either one or two vertices are ignored.
	PrimitiveTypeTriangle PrimitiveType = 3

	// PrimitiveTypeTriangleStrip rasterizes a triangle for every three adjacent vertices.
	PrimitiveTypeTriangleStrip PrimitiveType = 4
)

// LoadAction defines actions performed at the start of a rendering pass
// for a render command encoder.
//
// Reference: https://developer.apple.com/documentation/metal/mtlloadaction
type LoadAction uint8

// Actions performed at the start of a rendering pass for a render command encoder.
const (
	// LoadActionDontCare indicates that the GPU has permission to discard the existing contents of the attachment at the start of the render pass, replacing them with arbitrary data.
	LoadActionDontCare LoadAction = 0

	// LoadActionLoad indicates that the GPU preserves the existing contents of the attachment at the start of the render pass.
	LoadActionLoad LoadAction = 1

	// LoadActionClear indicates that the GPU writes a value to every pixel in the attachment at the start of the render pass.
	LoadActionClear LoadAction = 2
)

// StoreAction defines actions performed at the end of a rendering pass
// for a render command encoder.
//
// Reference: https://developer.apple.com/documentation/metal/mtlstoreaction
type StoreAction uint8

// Actions performed at the end of a rendering pass for a render command encoder.
const (
	// StoreActionDontCare indicates that the GPU has permission to discard the rendered contents of the attachment at the end of the render pass, replacing them with arbitrary data.
	StoreActionDontCare StoreAction = 0

	// StoreActionStore indicates that the GPU stores the rendered contents to the texture.
	StoreActionStore StoreAction = 1

	// StoreActionMultisampleResolve indicates that the GPU resolves the multisampled data to one sample per pixel and stores the data to the resolve texture, discarding the multisample data afterwards.
	StoreActionMultisampleResolve StoreAction = 2

	// StoreActionStoreAndMultisampleResolve indicates that the GPU stores the multisample data to the multisample texture, resolves the data to a sample per pixel, and stores the data to the resolve texture.
	StoreActionStoreAndMultisampleResolve StoreAction = 3

	// StoreActionUnknown indicates that the app will specify the store action when it encodes the render pass.
	StoreActionUnknown StoreAction = 4

	// StoreActionCustomSampleDepthStore indicates that the GPU stores depth data in a sample-position–agnostic representation.
	StoreActionCustomSampleDepthStore StoreAction = 5
)

// ClearColor is an RGBA value used for a color pixel.
//
// Reference: https://developer.apple.com/documentation/metal/mtlclearcolor
type ClearColor struct {
	Red, Green, Blue, Alpha float64
}

// Resource represents a memory allocation for storing specialized data
// that is accessible to the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtlresource
type Resource interface {
	// resource returns the underlying id<MTLResource> pointer.
	resource() unsafe.Pointer
}

// Size represents the set of dimensions that declare the size of an object,
// such as an image, texture, threadgroup, or grid.
//
// Reference: https://developer.apple.com/documentation/metal/mtlsize
type Size struct{ Width, Height, Depth uint }

// Origin represents the location of a pixel in an image or texture relative
// to the upper-left corner, whose coordinates are (0, 0).
//
// Reference: https://developer.apple.com/documentation/metal/mtlorigin
type Origin struct{ X, Y, Z uint }

// Region is a rectangular block of pixels in an image or texture,
// defined by its upper-left corner and its size.
//
// Reference: https://developer.apple.com/documentation/metal/mtlregion
type Region struct {
	Origin Origin // The location of the upper-left corner of the block.
	Size   Size   // The size of the block.
}

// RegionMake1D returns a 1D, rectangular region for image or texture data.
//
// Reference: https://developer.apple.com/documentation/metal/1515675-mtlregionmake1d
func RegionMake1D(x, width uint) Region {
	return Region{
		Origin: Origin{x, 0, 0},
		Size:   Size{width, 1, 1},
	}
}

// RegionMake2D returns a 2D, rectangular region for image or texture data.
//
// Reference: https://developer.apple.com/documentation/metal/1515675-mtlregionmake2d
func RegionMake2D(x, y, width, height uint) Region {
	return Region{
		Origin: Origin{x, y, 0},
		Size:   Size{width, height, 1},
	}
}

// RegionMake3D returns a 3D, rectangular region for image or texture data.
//
// Reference: https://developer.apple.com/documentation/metal/1515675-mtlregionmake3d
func RegionMake3D(x, y, z, width, height, depth uint) Region {
	return Region{
		Origin: Origin{x, y, z},
		Size:   Size{width, height, depth},
	}
}
