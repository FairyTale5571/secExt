package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"

import (
	"fmt"
	"net/http"
	"regexp"
	"runtime"
	"time"
	"unsafe"

	"github.com/fairytale5571/secExt/pkg/app"
	"github.com/fairytale5571/secExt/pkg/ds"
	"github.com/fairytale5571/secExt/pkg/env"
	"github.com/fairytale5571/secExt/pkg/files"
	"github.com/fairytale5571/secExt/pkg/ip"
)

var a *app.App

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
	PrintInArma(output, outputSize, "sec_v6")
}

//export goRVExtensionArgs
func goRVExtensionArgs(output *C.char, outputSize C.size_t, input *C.char, argv **C.char, argc C.int) C.int {
	var err error
	clearArgs := cleanInput(argv, int(argc))
	action := C.GoString(input)
	if a == nil {
		a, err = app.New()
		if err != nil {
			a.Logger.Fatalf("Error start application secExt: %s", err)
		}
	}
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
		PrintInArma(output, outputSize, a.CleanTemp())
	case "isAdmin":
		PrintInArma(output, outputSize, a.IsAdmin())
	case "get_HWID":
		r = a.ReadRegistry("local_machine", `SOFTWARE\Microsoft\Cryptography`, "MachineGuid")
		PrintInArma(output, outputSize, r)
	case "get_Process":
		PrintInArma(output, outputSize, a.GetProcesses())
	case "get_MAC":
		PrintInArma(output, outputSize, a.GetMacAddr())
	case "get_IP":
		PrintInArma(output, outputSize, ip.GetIp())
	case "get_GeoIP":
		PrintInArma(output, outputSize, ip.GetGeoIp())
	case "get_Sd":
		PrintInArma(output, outputSize, ds.GetDsName())
	case "v":
		PrintInArma(output, outputSize, runtime.Version())

	case "GetCPU_id":
		r, err = a.Wmi.GetCpuId()
		if err != nil {
			PrintInArma(output, outputSize, "cpu undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "GetCPU_name":
		r, err = a.Wmi.GetCpuName()
		if err != nil {
			PrintInArma(output, outputSize, "cpu undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "GetMother_id":
		r, err = a.Wmi.GetMotherId()
		if err != nil {
			PrintInArma(output, outputSize, "mother undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "GetMother_name":
		r, err = a.Wmi.GetMotherName()
		if err != nil {
			PrintInArma(output, outputSize, "mother undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "GetBios_id":
		r, err = a.Wmi.GetBiosId()
		if err != nil {
			PrintInArma(output, outputSize, "bios undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "GetBios_ReleaseDate":
		r, err = a.Wmi.GetBiosReleaseDate()
		if err != nil {
			PrintInArma(output, outputSize, "bios undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "GetBios_Version":
		r, err = a.Wmi.GetBiosVersion()
		if err != nil {
			PrintInArma(output, outputSize, "bios undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "GetRam_serialNumber":
		r, err = a.Wmi.GetRamSerialNumber()
		if err != nil {
			PrintInArma(output, outputSize, "ram undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "GetRam_capacity":
		r, err = a.Wmi.GetRamCapacity()
		if err != nil {
			PrintInArma(output, outputSize, "ram undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "GetRam_partNumber":
		r, err = a.Wmi.GetRamPartNumber()
		if err != nil {
			PrintInArma(output, outputSize, "ram undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "GetRam_Name":
		r, err = a.Wmi.GetRamName()
		if err != nil {
			PrintInArma(output, outputSize, "ram undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "GetProduct_Date":
		r, err = a.Wmi.GetProductInstallDate()
		if err != nil {
			PrintInArma(output, outputSize, "product undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "GetProduct_Version":
		r, err = a.Wmi.GetProductVersion()
		if err != nil {
			PrintInArma(output, outputSize, "product undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "Get_Drives":
		r, err = a.Wmi.GetDiskDrives()
		if err != nil {
			PrintInArma(output, outputSize, "drives undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "get_Product":
		r, err = a.Wmi.GetProductId()
		if err != nil {
			PrintInArma(output, outputSize, "product undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "GetPC_name":
		r, err = a.Wmi.GetPcName()
		if err != nil {
			PrintInArma(output, outputSize, "pc undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "Get_SID":
		r, err = a.Wmi.GetSID()
		if err != nil {
			PrintInArma(output, outputSize, fmt.Sprintf("sid undefined: %s", err.Error()))
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "Get_VRAM_name":
		r, err = a.Wmi.GetVRAM()
		if err != nil {
			PrintInArma(output, outputSize, "vram undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "get_CSP":
		r, err = a.Wmi.GetCSP()
		if err != nil {
			PrintInArma(output, outputSize, "CSP undefined")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "uuid":
		PrintInArma(output, outputSize, a.GetUUID())
	case "setEnv":
		if len(clearArgs) < 2 {
			a.Logger.Errorf("setEnv: not enough args")
			PrintInArma(output, outputSize, "setEnv: not enough arguments")
			return http.StatusConflict
		}
		go env.Set(clearArgs[0], clearArgs[1])
		PrintInArma(output, outputSize, "ok")
	case "getEnv":
		if len(clearArgs) < 1 {
			a.Logger.Errorf("getEnv: not enough arguments")
			PrintInArma(output, outputSize, "getEnv: not enough arguments")
			return http.StatusConflict
		}
		PrintInArma(output, outputSize, env.Get(clearArgs[0]))
	case "1_c":
		r, err = a.Drive.SetCredentials(clearArgs[0])
		if err != nil {
			a.Logger.Errorf("set credentials: %s", err.Error())
			PrintInArma(output, outputSize, "error")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "2_c":
		r, err = a.Drive.SetToken(clearArgs[0])
		if err != nil {
			a.Logger.Errorf("set credentials: %s", err.Error())
			PrintInArma(output, outputSize, "error")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "3_c":
		time.Sleep(5 * time.Second)
		if err := a.Drive.DumpScreen(fmt.Sprintf("/%s_%s", clearArgs[0], clearArgs[1])); err != nil {
			a.Logger.Errorf("dump screen: %s", err.Error())
			PrintInArma(output, outputSize, "error")
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, "ok")
	case "3_c_t":
		go func() {
			if err := a.Drive.DumpScreen(fmt.Sprintf("/%s_%s", clearArgs[0], clearArgs[1])); err != nil {
				a.Logger.Errorf("dump screen: %s", err.Error())
				PrintInArma(output, outputSize, "error")
				return
			}
		}()
		PrintInArma(output, outputSize, "ok")
	case "1_r":
		if len(clearArgs) < 4 {
			a.Logger.Errorf("1_r: not enough arguments")
			PrintInArma(output, outputSize, "1_r: not enough arguments")
			return http.StatusConflict
		}
		err = a.Reg.WriteReg(clearArgs[0], clearArgs[1], clearArgs[2], clearArgs[3])
		if err != nil {
			PrintInArma(output, outputSize, fmt.Sprintf("writeReg: %s", err.Error()))
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, "Written")
	case "2_r":
		if len(clearArgs) < 3 {
			a.Logger.Errorf("2_r: not enough arguments")
			PrintInArma(output, outputSize, "2_r: not enough arguments")
			return http.StatusConflict
		}
		r, err = a.Reg.ReadReg(clearArgs[0], clearArgs[1], clearArgs[2])
		if err != nil {
			PrintInArma(output, outputSize, fmt.Sprintf("reg undefined: %s", err.Error()))
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "3_r":
		if len(clearArgs) < 3 {
			a.Logger.Errorf("3_r: not enough arguments")
			PrintInArma(output, outputSize, "3_r: not enough arguments")
			return http.StatusConflict
		}
		err = a.Reg.DeleteReg(clearArgs[0], clearArgs[1], clearArgs[2])
		if err != nil {
			PrintInArma(output, outputSize, fmt.Sprintf("deleteReg: %s", err.Error()))
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, "Deleted")
	case "1_f":
		if len(clearArgs) < 2 {
			a.Logger.Errorf("1_f: not enough arguments")
			PrintInArma(output, outputSize, "1_f: not enough arguments")
			return http.StatusConflict
		}
		err = files.WriteFile(clearArgs[0], clearArgs[1])
		if err != nil {
			PrintInArma(output, outputSize, fmt.Sprintf("writeFile: %s", err.Error()))
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, "Written")
	case "2_f":
		if len(clearArgs) < 1 {
			a.Logger.Errorf("2_f: not enough arguments")
			PrintInArma(output, outputSize, "2_f: not enough arguments")
			return http.StatusConflict
		}
		r, err = files.ReadFile(clearArgs[0])
		if err != nil {
			PrintInArma(output, outputSize, fmt.Sprintf("readFile: %s", err.Error()))
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, r)
	case "3_f":
		if len(clearArgs) < 1 {
			a.Logger.Errorf("3_f: not enough arguments")
			PrintInArma(output, outputSize, "3_f: not enough arguments")
			return http.StatusConflict
		}
		err = files.DeleteFile(clearArgs[0])
		if err != nil {
			PrintInArma(output, outputSize, fmt.Sprintf("deleteFile: %s", err.Error()))
			return http.StatusInternalServerError
		}
		PrintInArma(output, outputSize, "Deleted")
	default:
		temp := fmt.Sprintf("Function: %s nb params: %d params: %s!", C.GoString(input), argc, clearArgs)
		a.Logger.Errorf(temp)
		PrintInArma(output, outputSize, temp)
		return http.StatusNotFound
	}
	return http.StatusOK
}

func PrintInArma(output *C.char, outputSize C.size_t, input string) {
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
	_ = goRVExtensionArgs(output, outputSize, input, nil, C.int(0))
}

func main() {}
