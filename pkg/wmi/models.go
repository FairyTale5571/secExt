package wmi

import (
	"fmt"
	"time"

	"github.com/StackExchange/wmi"
)

type Win32Processor struct {
	ProcessorId string
	Name        string
	SystemName  string
}

func getCPU() []Win32Processor {
	var dst []Win32Processor
	if err := wmi.Query("SELECT * FROM Win32_Processor", &dst); err != nil {
		// TODO: logger
	}
	return dst
}

type Win32BaseBoard struct {
	Product      string
	SerialNumber string
}

func getMother() []Win32BaseBoard {
	var dst []Win32BaseBoard
	if err := wmi.Query("SELECT * FROM Win32_BaseBoard", &dst); err != nil {
		// TODO: logger
	}
	return dst
}

type Win32BIOS struct {
	SerialNumber string
	ReleaseDate  time.Time
	Version      string
}

func getBios() []Win32BIOS {
	var dst []Win32BIOS
	if err := wmi.Query("SELECT * FROM Win32_BIOS", &dst); err != nil {
		// TODO: logger
	}
	return dst
}

type Win32PhysicalMemory struct {
	SerialNumber string
	PartNumber   string
	Manufacturer string
	Capacity     uint64
}

func getRAM() []Win32PhysicalMemory {
	var dst []Win32PhysicalMemory
	if err := wmi.Query("SELECT * FROM Win32_PhysicalMemory", &dst); err != nil {
		// TODO: logger
	}
	return dst
}

type Win32OperatingSystem struct {
	Version      string
	InstallDate  time.Time
	SerialNumber string
}

func getOS() []Win32OperatingSystem {
	var dst []Win32OperatingSystem
	if err := wmi.Query("SELECT * FROM Win32_OperatingSystem", &dst); err != nil {
		// TODO: logger
	}
	return dst
}

type Win32ComputerSystemProduct struct {
	Caption           string
	Description       string
	IdentifyingNumber string
	Name              string
	SKUNumber         string
	Vendor            string
	Version           string
	UUID              string
}

func getCSP() []Win32ComputerSystemProduct {
	var dst []Win32ComputerSystemProduct
	if err := wmi.Query("SELECT * FROM Win32_ComputerSystemProduct", &dst); err != nil {
		// TODO: logger
	}
	return dst
}

type Win32VideoController struct {
	Name string
}

func getVRAM() string {
	var dst []Win32VideoController
	if err := wmi.Query("SELECT * FROM Win32_VideoController", &dst); err != nil {
		// TODO: logger
	}

	return fmt.Sprintf("GPU %s", dst[0].Name)
}

type Win32DiskDrive struct {
	Model        string
	SerialNumber string
	Size         uint64
}

func getDiskDrive() []Win32DiskDrive {
	var dst []Win32DiskDrive
	if err := wmi.Query("SELECT * FROM Win32_DiskDrive", &dst); err != nil {
		// TODO: logger
	}
	return dst
}

type Win32UserAccount struct {
	SID string
}

func getUserAccount() *Win32UserAccount {
	var dst Win32UserAccount
	if err := wmi.Query("SELECT * FROM Win32_UserAccount ", &dst); err != nil {
		// TODO: logger
	}
	return &dst
}
