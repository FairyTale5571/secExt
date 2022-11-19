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

func (w *Wmi) GetCpuId() (string, error) {
	r, err := getCPU()
	if err != nil {
		return "", nil
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no cpu found")
	}
	return fmt.Sprintf(r[0].ProcessorId), nil
}

func (w *Wmi) GetCpuName() (string, error) {
	r, err := getCPU()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no cpu found")
	}
	return fmt.Sprintf(r[0].Name), nil
}

func (w *Wmi) GetMotherId() (string, error) {
	r, err := getMother()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no motherboard found")
	}
	return fmt.Sprintf(r[0].SerialNumber), nil
}

func (w *Wmi) GetMotherName() (string, error) {
	r, err := getMother()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no motherboard found")
	}
	return fmt.Sprintf("Motherboard %s", r[0].Product), nil
}

func (w *Wmi) GetRamSerialNumber() (string, error) {
	var numbers []string
	r, err := getRAM()
	if err != nil {
		return "", err
	}
	for idx := range r {
		numbers = append(numbers, r[idx].SerialNumber)
	}
	return helpers.Struct2JSON(numbers), nil
}

func (w *Wmi) GetRamPartNumber() (string, error) {
	var numbers []string
	r, err := getRAM()
	if err != nil {
		return "", err
	}

	for _, v := range r {
		numbers = append(numbers, v.PartNumber)
	}
	return helpers.Struct2JSON(numbers), nil
}

func (w *Wmi) GetRamName() (string, error) {
	var numbers []string
	r, err := getRAM()
	if err != nil {
		return "", err
	}

	for _, v := range r {
		numbers = append(numbers, v.Manufacturer)
	}
	return helpers.Struct2JSON(numbers), nil
}

func (w *Wmi) GetRamCapacity() (string, error) {
	var memory uint64
	r, err := getRAM()
	if err != nil {
		return "", err
	}
	for _, v := range r {
		memory += v.Capacity
	}
	size, bytef := helpers.ConvertSize(memory)
	return fmt.Sprintf("%d %v", size, bytef), nil
}

func (w *Wmi) GetProductId() (string, error) {
	r, err := getOS()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no os found")
	}
	return fmt.Sprintf(r[0].SerialNumber), nil
}

func (w *Wmi) GetProductInstallDate() (string, error) {
	r, err := getOS()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no os found")
	}
	return fmt.Sprintf("%v", r[0].InstallDate), nil
}

func (w *Wmi) GetProductVersion() (string, error) {
	r, err := getOS()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no os found")
	}
	return fmt.Sprintf(r[0].Version), nil
}

func (w *Wmi) GetBiosId() (string, error) {
	r, err := getBios()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no bios found")
	}

	return fmt.Sprintf(r[0].SerialNumber), nil
}

func (w *Wmi) GetBiosReleaseDate() (string, error) {
	r, err := getBios()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no bios found")
	}
	return r[0].ReleaseDate.String(), nil
}

func (w *Wmi) GetBiosVersion() (string, error) {
	r, err := getBios()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no bios found")
	}
	return fmt.Sprintf(r[0].Version), nil
}

func (w *Wmi) GetPcName() (string, error) {
	r, err := getCPU()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no cpu found")
	}
	return fmt.Sprintf(r[0].SystemName), nil
}

func (w *Wmi) GetSID() (string, error) {
	r, err := getUserAccount()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("sid array empty")
	}
	return fmt.Sprintf(r[0].SID), nil
}

func (w *Wmi) GetVRAM() (string, error) {
	r, err := getVRAM()
	if err != nil {
		return "", err
	}
	return r, nil
}

func (w *Wmi) GetCSP() (string, error) {
	r, err := getCSP()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no csp found")
	}
	return r[0].Name, nil
}

func (w *Wmi) GetDiskDrives() (string, error) {
	drives, err := getDiskDrive()
	if err != nil {
		return "", err
	}
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
	return drive, nil
}
