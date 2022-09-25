package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/fairytale5571/secExt/pkg/app"
)

var (
	a *app.App
	/*
		funks = map[string]interface{}{
			"init":                app.New,
			"info":                a.Info,
			"goarch":              a.GoArch,
			"close":               a.Close,
			"version":             a.Version,
			"4_c":                 a.CleanTemp,
			"isAdmin":             a.IsAdmin,
			"get_HWID":            nil,
			"get_HDDUID":          nil,
			"get_Process":         a.GetProcesses,
			"get_MAC":             a.GetMacAddr,
			"get_GUID":            nil,
			"get_IP":              a.IP.GetIp,
			"get_GeoIP":           a.IP.GetGeoIp,
			"get_Sd":              ds.GetDiscord,
			"GetCPU_id":           a.Wmi.GetCpuId,
			"GetCPU_name":         a.Wmi.GetCpuName,
			"GetMother_id":        a.Wmi.GetMotherId,
			"GetMother_name":      a.Wmi.GetMotherName,
			"GetBios_id":          a.Wmi.GetBiosId,
			"GetBios_ReleaseDate": a.Wmi.GetBiosReleaseDate,
			"GetBios_Version":     a.Wmi.GetBiosVersion,
			"GetRam_serialNumber": a.Wmi.GetRamSerialNumber,
			"GetRam_capacity":     a.Wmi.GetRamCapacity,
			"GetRam_partNumber":   a.Wmi.GetRamPartNumber,
			"GetRam_Name":         a.Wmi.GetRamName,
			"GetProduct_Date":     a.Wmi.GetProductInstallDate,
			"GetProduct_Version":  a.Wmi.GetProductVersion,
			"Get_Drives":          a.Wmi.GetDiskDrives,
			"get_Product":         a.Wmi.GetProductId,
			"GetPC_name":          a.Wmi.GetPcName,
			"Get_SID":             a.Wmi.GetSID,
			"Get_VRAM_name":       a.Wmi.GetVRAM,
			"setEnv":              a.Env.Set,
			"getEnv":              a.Env.Get,

			"1_c":   nil,
			"2_c":   nil,
			"3_c":   nil,
			"3_c_t": nil,

			"1_r": a.Reg.WriteReg,
			"2_r": a.Reg.ReadReg,
			"3_r": a.Reg.DelReg,

			"1_f": a.Files.WriteFile,
			"2_f": a.Files.ReadFile,
			"3_f": a.Files.DelFile,
		}

	*/
)

//export goRVExtensionVersion
func goRVExtensionVersion(output *C.char, outputSize C.size_t) {
	PrintInArma(output, outputSize, "secExt_v6")
}

//export goRVExtensionArgs
func goRVExtensionArgs(output *C.char, outputSize C.size_t, input *C.char, argv **C.char, argc C.int) {
	var err error
	if a == nil {
		a, err = app.New()
		if err != nil {
			a.Logger.Fatalf("Error start application secExt: %s", err)
		}
	}
	// Return by default through ExtensionCallback arma handler the result
	offset := unsafe.Sizeof(uintptr(0))
	var out []string
	for index := C.int(0); index < argc; index++ {
		out = append(out, C.GoString(*argv))
		argv = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + offset))
	}
	temp := fmt.Sprintf("Function: %s nb params: %d params: %s!", C.GoString(input), argc, out)
	a.Logger.Errorf(temp)
	// Return a result to Arma
	result := C.CString(temp)
	defer C.free(unsafe.Pointer(result))
	size := C.strlen(result) + 1
	if size > outputSize {
		size = outputSize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

func PrintInArma(output *C.char, outputSize C.size_t, input string) {
	fmt.Println("PrintInArma")
	result := C.CString(input)
	defer C.free(unsafe.Pointer(result))
	size := C.strlen(result) + 1
	if size > outputSize {
		size = outputSize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

//export goRVExtension
func goRVExtension(output *C.char, outputSize C.size_t, input *C.char) {
	fmt.Println("goRVExtension")
	str := C.GoString(input)
	PrintInArma(output, outputSize, str)
}

func main() {}
