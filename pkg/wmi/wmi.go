package wmi

import (
	"fmt"
	"os"

	"github.com/fairytale5571/secExt/pkg/cache"
	"github.com/fairytale5571/secExt/pkg/helpers"
)

type Wmi struct {
	cache *cache.Config
}

func New() *Wmi {
	return &Wmi{
		cache: cache.SetupCache(),
	}
}

func (w *Wmi) GetCpuId() (string, error) {
	if w.cache.IsExist("cpu_id") {
		return w.cache.Get("cpu_id")
	}

	r, err := getCPU()
	if err != nil {
		return "", nil
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no cpu found")
	}
	w.cache.Set("cpu_id", fmt.Sprintf(r[0].ProcessorId))
	return fmt.Sprintf(r[0].ProcessorId), nil
}

func (w *Wmi) GetCpuName() (string, error) {
	if w.cache.IsExist("cpu_name") {
		return w.cache.Get("cpu_name")
	}

	r, err := getCPU()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no cpu found")
	}
	w.cache.Set("cpu_name", fmt.Sprintf(r[0].Name))
	return fmt.Sprintf(r[0].Name), nil
}

func (w *Wmi) GetMotherId() (string, error) {
	if w.cache.IsExist("mother_id") {
		return w.cache.Get("mother_id")
	}

	r, err := getMother()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no motherboard found")
	}
	w.cache.Set("mother_id", fmt.Sprintf(r[0].SerialNumber))
	return fmt.Sprintf(r[0].SerialNumber), nil
}

func (w *Wmi) GetMotherName() (string, error) {
	if w.cache.IsExist("mother_name") {
		return w.cache.Get("mother_name")
	}

	r, err := getMother()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no motherboard found")
	}
	w.cache.Set("mother_name", fmt.Sprintf(r[0].Product))
	return fmt.Sprintf("Motherboard %s", r[0].Product), nil
}

func (w *Wmi) GetRamSerialNumber() (string, error) {
	if w.cache.IsExist("ram_serial_number") {
		return w.cache.Get("ram_serial_number")
	}

	var numbers []string
	r, err := getRAM()
	if err != nil {
		return "", err
	}
	for idx := range r {
		numbers = append(numbers, r[idx].SerialNumber)
	}
	res := helpers.Struct2JSON(numbers)
	w.cache.Set("ram_serial_number", res)
	return res, nil
}

func (w *Wmi) GetRamPartNumber() (string, error) {
	if w.cache.IsExist("ram_part_number") {
		return w.cache.Get("ram_part_number")
	}

	var numbers []string
	r, err := getRAM()
	if err != nil {
		return "", err
	}

	for _, v := range r {
		numbers = append(numbers, v.PartNumber)
	}
	res := helpers.Struct2JSON(numbers)
	w.cache.Set("ram_part_number", res)
	return res, nil
}

func (w *Wmi) GetRamName() (string, error) {
	if w.cache.IsExist("ram_name") {
		return w.cache.Get("ram_name")
	}

	var numbers []string
	r, err := getRAM()
	if err != nil {
		return "", err
	}

	for _, v := range r {
		numbers = append(numbers, v.Manufacturer)
	}
	res := helpers.Struct2JSON(numbers)
	w.cache.Set("ram_name", res)
	return res, nil
}

func (w *Wmi) GetRamCapacity() (string, error) {
	if w.cache.IsExist("ram_capacity") {
		return w.cache.Get("ram_capacity")
	}
	var memory uint64
	r, err := getRAM()
	if err != nil {
		return "", err
	}
	for _, v := range r {
		memory += v.Capacity
	}
	size, bytef := helpers.ConvertSize(memory)
	res := fmt.Sprintf("%v %s", size, bytef)
	w.cache.Set("ram_capacity", res)
	return res, nil
}

func (w *Wmi) GetProductId() (string, error) {
	if w.cache.IsExist("product_id") {
		return w.cache.Get("product_id")
	}
	r, err := getOS()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no os found")
	}
	w.cache.Set("product_id", r[0].SerialNumber)
	return r[0].SerialNumber, nil
}

func (w *Wmi) GetProductInstallDate() (string, error) {
	if w.cache.IsExist("product_install_date") {
		return w.cache.Get("product_install_date")
	}

	r, err := getOS()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no os found")
	}

	w.cache.Set("product_install_date", fmt.Sprintf("%v", r[0].InstallDate))
	return fmt.Sprintf("%v", r[0].InstallDate), nil
}

func (w *Wmi) GetProductVersion() (string, error) {
	if w.cache.IsExist("product_version") {
		return w.cache.Get("product_version")
	}

	r, err := getOS()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no os found")
	}
	w.cache.Set("product_version", r[0].Version)
	return fmt.Sprintf(r[0].Version), nil
}

func (w *Wmi) GetBiosId() (string, error) {
	if w.cache.IsExist("bios_id") {
		return w.cache.Get("bios_id")
	}
	r, err := getBios()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no bios found")
	}

	w.cache.Set("bios_id", r[0].SerialNumber)
	return fmt.Sprintf(r[0].SerialNumber), nil
}

func (w *Wmi) GetBiosReleaseDate() (string, error) {
	if w.cache.IsExist("bios_release_date") {
		return w.cache.Get("bios_release_date")
	}

	r, err := getBios()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no bios found")
	}
	w.cache.Set("bios_release_date", r[0].ReleaseDate.String())
	return r[0].ReleaseDate.String(), nil
}

func (w *Wmi) GetBiosVersion() (string, error) {
	if w.cache.IsExist("bios_version") {
		return w.cache.Get("bios_version")
	}

	r, err := getBios()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no bios found")
	}
	w.cache.Set("bios_version", r[0].Version)
	return fmt.Sprintf(r[0].Version), nil
}

func (w *Wmi) GetPcName() (string, error) {
	if w.cache.IsExist("pc_name") {
		return w.cache.Get("pc_name")
	}
	r, err := os.Hostname()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no cpu found")
	}
	w.cache.Set("pc_name", r)
	return r, nil
}

func (w *Wmi) GetSID() (string, error) {
	if w.cache.IsExist("sid") {
		return w.cache.Get("sid")
	}
	r, err := getUserAccount()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("sid array empty")
	}
	w.cache.Set("sid", r[0].SID)
	return fmt.Sprintf(r[0].SID), nil
}

func (w *Wmi) GetVRAM() (string, error) {
	if w.cache.IsExist("vram") {
		return w.cache.Get("vram")
	}
	r, err := getVRAM()
	if err != nil {
		return "", err
	}

	w.cache.Set("vram", r)
	return r, nil
}

func (w *Wmi) GetCSP() (string, error) {
	if w.cache.IsExist("csp") {
		return w.cache.Get("csp")
	}

	r, err := getCSP()
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", fmt.Errorf("no csp found")
	}

	w.cache.Set("csp", r[0].Name)
	return r[0].Name, nil
}

func (w *Wmi) GetDiskDrives() (string, error) {
	if w.cache.IsExist("disk_drives") {
		return w.cache.Get("disk_drives")
	}

	drives, err := getDiskDrive()
	if err != nil {
		return "", err
	}
	drive := "["
	var size uint64
	var str string
	for idx := range drives {
		size, str = helpers.ConvertSize(drives[idx].Size)
		drive += fmt.Sprintf(`["%v %d %v %v"],`, drives[idx].Model, size, str, drives[idx].SerialNumber)
	}
	drive = drive[:len(drive)-1]
	drive += "]"
	w.cache.Set("disk_drives", drive)
	return drive, nil
}
