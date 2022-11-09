package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"
import (
	"fmt"
	"regexp"
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

func cleanInput(argv **C.char, argc int) []string {
	newArgs := make([]string, argc)
	offset := unsafe.Sizeof(uintptr(0))
	i := 0
	for i < argc {
		_arg := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + offset*uintptr(i)))
		arg := C.GoString(*_arg)
		arg = arg[1 : len(arg)-1]

		reArg := regexp.MustCompile(`""`)
		arg = reArg.ReplaceAllString(arg, `"`)

		newArgs[i] = arg
		i++
	}
	return newArgs
}

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
	clearArgs := cleanInput(argv, int(argc))
	action := C.GoString(input)
	var r string
	switch action {
	case "info":
		PrintInArma(output, outputSize, a.Info())
	case "goarch":
		PrintInArma(output, outputSize, a.GoArch())
	case "close":
		a.Close()
	case "version":
		PrintInArma(output, outputSize, a.Version())
	case "4_c":
		a.CleanTemp()
	case "isAdmin":
		PrintInArma(output, outputSize, a.IsAdmin())
	case "get_HWID":
		r = a.ReadRegistry("HKEY_LOCAL_MACHINE", `SOFTWARE\Microsoft\Cryptography`, "MachineGuid")
		PrintInArma(output, outputSize, r)
	case "get_HDDUID":
		PrintInArma(output, outputSize, "deprecated")
	case "get_Process":
		PrintInArma(output, outputSize, a.GetProcesses())
	case "get_MAC":
		PrintInArma(output, outputSize, a.GetMacAddr())
	case "get_GUID":
		r = a.ReadRegistry("current_user", `Software\Classes\mscfile\shell\open\command`, "GUID")
		PrintInArma(output, outputSize, r)
	case "get_IP":
		PrintInArma(output, outputSize, a.GetIP())
	case "get_GeoIP":
		PrintInArma(output, outputSize, a.GetGeoIP())
	case "get_Sd":
		PrintInArma(output, outputSize, a.GetDiscordArray())
	case "get_SD_ID":
		PrintInArma(output, outputSize, a.GetDiscordID())
	case "get_SD_User":
		PrintInArma(output, outputSize, a.GetDiscordUsername())
	case "GetCPU_id":
		r, err = a.Wmi.GetCpuId()
		if err != nil {
			PrintInArma(output, outputSize, "cpu undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "GetCPU_name":
		r, err = a.Wmi.GetCpuName()
		if err != nil {
			PrintInArma(output, outputSize, "cpu undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "GetMother_id":
		r, err = a.Wmi.GetMotherId()
		if err != nil {
			PrintInArma(output, outputSize, "mother undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "GetMother_name":
		r, err = a.Wmi.GetMotherName()
		if err != nil {
			PrintInArma(output, outputSize, "mother undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "GetBios_id":
		r, err = a.Wmi.GetBiosId()
		if err != nil {
			PrintInArma(output, outputSize, "bios undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "GetBios_ReleaseDate":
		r, err = a.Wmi.GetBiosReleaseDate()
		if err != nil {
			PrintInArma(output, outputSize, "bios undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "GetBios_Version":
		r, err = a.Wmi.GetBiosVersion()
		if err != nil {
			PrintInArma(output, outputSize, "bios undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "GetRam_serialNumber":
		r, err = a.Wmi.GetRamSerialNumber()
		if err != nil {
			PrintInArma(output, outputSize, "ram undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "GetRam_capacity":
		r, err = a.Wmi.GetRamCapacity()
		if err != nil {
			PrintInArma(output, outputSize, "ram undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "GetRam_partNumber":
		r, err = a.Wmi.GetRamPartNumber()
		if err != nil {
			PrintInArma(output, outputSize, "ram undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "GetRam_Name":
		r, err = a.Wmi.GetRamName()
		if err != nil {
			PrintInArma(output, outputSize, "ram undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "GetProduct_Date":
		r, err = a.Wmi.GetProductInstallDate()
		if err != nil {
			PrintInArma(output, outputSize, "product undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "GetProduct_Version":
		r, err = a.Wmi.GetProductVersion()
		if err != nil {
			PrintInArma(output, outputSize, "product undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "Get_Drives":
		r, err = a.Wmi.GetDiskDrives()
		if err != nil {
			PrintInArma(output, outputSize, "drives undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "get_Product":
		r, err = a.Wmi.GetProductId()
		if err != nil {
			PrintInArma(output, outputSize, "product undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "GetPC_name":
		r, err = a.Wmi.GetPcName()
		if err != nil {
			PrintInArma(output, outputSize, "pc undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "Get_SID":
		r, err = a.Wmi.GetSID()
		if err != nil {
			PrintInArma(output, outputSize, "sid undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "Get_VRAM_name":
		r, err = a.Wmi.GetVRAM()
		if err != nil {
			PrintInArma(output, outputSize, "vram undefined")
			return
		}
		PrintInArma(output, outputSize, r)
	case "setEnv":
		if len(clearArgs) < 2 {
			PrintInArma(output, outputSize, "setEnv: not enough arguments")
			return
		}
		go a.SetEnv(clearArgs[0], clearArgs[1])
		PrintInArma(output, outputSize, "ok")
	case "getEnv":
		if len(clearArgs) < 1 {
			PrintInArma(output, outputSize, "getEnv: not enough arguments")
			return
		}
		PrintInArma(output, outputSize, a.GetEnv(clearArgs[0]))
	case "1_c":
	case "2_c":
	case "3_c":
	case "3_c_t":
	case "1_r":
	case "2_r":
	case "3_r":
	case "1_f":
	case "2_f":
	case "3_f":
	default:
		temp := fmt.Sprintf("Function: %s nb params: %d params: %s!", C.GoString(input), argc, clearArgs)
		a.Logger.Errorf(temp)
		PrintInArma(output, outputSize, temp)
	}
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
	goRVExtensionArgs(output, outputSize, input, nil, C.int(0))
}

func main() {}
