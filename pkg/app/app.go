package app

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"runtime"

	"github.com/fairytale5571/secExt/pkg/env"
	"github.com/fairytale5571/secExt/pkg/files"
	"github.com/fairytale5571/secExt/pkg/helpers"
	"github.com/fairytale5571/secExt/pkg/ip"
	"github.com/fairytale5571/secExt/pkg/logger"
	"github.com/fairytale5571/secExt/pkg/reg"
	"github.com/fairytale5571/secExt/pkg/wmi"
	"github.com/mitchellh/go-ps"
)

type App struct {
	Logger *logger.Wrapper
	Wmi    *wmi.Wmi
	Reg    *reg.Reg
	Files  *files.Files
	IP     *ip.IP
	Env    *env.Env
}

func New(args ...string) (*App, error) {
	log := logger.New("app")
	return &App{
		Logger: log,
		Wmi:    wmi.New(),
		Reg:    reg.New(),
		Files:  files.New(),
		IP:     ip.New(),
		Env:    env.New(),
	}, nil
}

func (a *App) Info() string {
	return "Extension developed by FairyTale#5571"
}

func (a *App) Version() string {
	return "1.0:553544"
}

func (a *App) GoArch() string {
	return runtime.GOARCH
}

func (a *App) IsAdmin() string {
	return helpers.IsAdmin()
}

func (a *App) CleanTemp() string {
	path := os.TempDir() + "/chrome_drag0947_254420441/dir/"
	err := os.RemoveAll(path)
	if err != nil {
		a.Logger.Errorf("error remove temp dir: %v", err)
		return "false"
	}
	return "true"
}

func (a *App) Close() string {
	os.Exit(1)
	return "Closing..."
}

func (a *App) GetMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && !bytes.Equal(i.HardwareAddr, nil) {
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}

func (a *App) GetProcesses() string {
	procs, err := ps.Processes()
	if err != nil {
		a.Logger.Errorf("error get processes: %v", err)
		return "error"
	}

	result := make(map[string]struct{})
	for _, proc := range procs {
		name := proc.Executable()
		if _, ok := result[name]; !ok {
			result[name] = struct{}{}
		}
	}
	var names []string
	for key := range result {
		names = append(names, key)
	}
	return fmt.Sprintf("%v\n", helpers.Struct2JSON(names))
}
