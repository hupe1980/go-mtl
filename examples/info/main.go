package main

import (
	"fmt"
	"log"

	"github.com/hupe1980/go-mtl"
)

func main() {
	device, err := mtl.CreateSystemDefaultDevice()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Preferred system default Metal device:", device.Name)

	allDevices := mtl.CopyAllDevices()

	for _, d := range allDevices {
		fmt.Println()
		printDeviceInfo(d)
	}
}

func printDeviceInfo(d mtl.Device) {
	fmt.Println(d.Name + ":")
	fmt.Println("• Low-power:", d.LowPower)
	fmt.Println("• Removable:", d.Removable)
	fmt.Println("• Configured as headless:", d.Headless)
	fmt.Println("• Registry ID:", d.RegistryID)
	fmt.Println()
	fmt.Println("Supports GPU Families:")
	fmt.Println("• Apple family 1 GPU features:", supported(d.SupportsFamily(mtl.GPUFamilyApple1)))
	fmt.Println("• Apple family 2 GPU features:", supported(d.SupportsFamily(mtl.GPUFamilyApple2)))
	fmt.Println("• Apple family 3 GPU features:", supported(d.SupportsFamily(mtl.GPUFamilyApple3)))
	fmt.Println("• Apple family 4 GPU features:", supported(d.SupportsFamily(mtl.GPUFamilyApple4)))
	fmt.Println("• Apple family 5 GPU features:", supported(d.SupportsFamily(mtl.GPUFamilyApple5)))
	fmt.Println("• Apple family 6 GPU features:", supported(d.SupportsFamily(mtl.GPUFamilyApple6)))
	fmt.Println("• Apple family 7 GPU features:", supported(d.SupportsFamily(mtl.GPUFamilyApple7)))
	fmt.Println("• Apple family 8 GPU features:", supported(d.SupportsFamily(mtl.GPUFamilyApple8)))
	fmt.Println("• Mac family 2 GPU features:", supported(d.SupportsFamily(mtl.GPUFamilyMac2)))
	fmt.Println("• Common family 1 GPU features:", supported(d.SupportsFamily(mtl.GPUFamilyCommon1)))
	fmt.Println("• Common family 2 GPU features:", supported(d.SupportsFamily(mtl.GPUFamilyCommon2)))
	fmt.Println("• Common family 3 GPU features:", supported(d.SupportsFamily(mtl.GPUFamilyCommon3)))
	fmt.Println("• Metal 3 features:", supported(d.SupportsFamily(mtl.GPUFamilyMetal3)))
}

func supported(v bool) string {
	switch v {
	case true:
		return "✅ supported"
	case false:
		return "❌ unsupported"
	}

	panic("unreachable")
}
