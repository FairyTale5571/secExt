package wmi

import (
	"fmt"

	"github.com/fairytale5571/secExt/pkg/helpers"
	"github.com/fairytale5571/secExt/pkg/logger"
)

type Wmi struct {
	logger *logger.Wrapper
}

func New() *Wmi {
	return &Wmi{
		logger: logger.New("wmi"),
	}
}

func (w *Wmi) GetCpuId() string {
	return fmt.Sprintf(getCPU()[0].ProcessorId)
}

func (w *Wmi) GetCpuName() string {
	return fmt.Sprintf(getCPU()[0].Name)
}

func (w *Wmi) GetMotherId() string {
	return fmt.Sprintf(getMother()[0].SerialNumber)
}

func (w *Wmi) GetMotherName() string {
	return fmt.Sprintf("Motherboard %s", getMother()[0].Product)
}

func (w *Wmi) GetRamSerialNumber() string {
	var numbers []string

	for idx := range getRAM() {
		numbers = append(numbers, getRAM()[idx].SerialNumber)
	}
	return helpers.Struct2JSON(numbers)
}

func (w *Wmi) GetRamPartNumber() string {
	var numbers []string

	for idx := range getRAM() {
		numbers = append(numbers, getRAM()[idx].PartNumber)
	}
	return helpers.Struct2JSON(numbers)
}

func (w *Wmi) GetRamName() string {
	var numbers []string

	for idx := range getRAM() {
		numbers = append(numbers, getRAM()[idx].Manufacturer)
	}
	return helpers.Struct2JSON(numbers)
}

func (w *Wmi) GetRamCapacity() string {
	var memory uint64 = 0
	for idx := range getRAM() {
		memory += getRAM()[idx].Capacity
	}
	size, bytef := helpers.ConvertSize(memory)
	return fmt.Sprintf("%d %v", size, bytef)
}

func (w *Wmi) GetProductId() string {
	return fmt.Sprintf(getOS()[0].SerialNumber)
}

func (w *Wmi) GetProductInstallDate() string {
	return fmt.Sprintf("%v", getOS()[0].InstallDate)
}

func (w *Wmi) GetProductVersion() string {
	return fmt.Sprintf(getOS()[0].Version)
}

func (w *Wmi) GetBiosId() string {
	return fmt.Sprintf(getBios()[0].SerialNumber)
}

func (w *Wmi) GetBiosReleaseDate() string {
	return getBios()[0].ReleaseDate.String()
}

func (w *Wmi) GetBiosVersion() string {
	return fmt.Sprintf(getBios()[0].Version)
}

func (w *Wmi) GetPcName() string {
	return fmt.Sprintf(getCPU()[0].SystemName)
}

func (w *Wmi) GetSID() string {
	return fmt.Sprintf(getUserAccount().SID)
}

func (w *Wmi) GetVRAM() string {
	return getVRAM()
}

func (w *Wmi) GetCSP() string {
	return getCSP()[0].Name
}

func (w *Wmi) GetDiskDrives() string {
	drives := getDiskDrive()

	drive := "["
	var size uint64
	var str string
	elems := len(drives)
	for idx := range drives {
		size, str = helpers.ConvertSize(drives[idx].Size)
		drive += fmt.Sprintf(`["%v %d %v %v"]`, drives[idx].Model, size, str, drives[idx].SerialNumber)
		if elems-1 != idx {
			drive += ","
		}
	}
	drive += "]"
	return drive
}
