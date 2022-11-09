package app

import (
	"bytes"
	"fmt"
	"github.com/fairytale5571/secExt/pkg/ds"
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
	Logger  *logger.Wrapper
	Wmi     *wmi.Wmi
	Reg     *reg.Reg
	Files   *files.Files
	IP      *ip.IP
	Env     *env.Env
	discord *ds.DS
}

func New(args ...string) (*App, error) {
	log := logger.New("app")
	dsRPC, err := ds.New()
	if err != nil {
		log.Errorf("error init discord: %v", err)
		return nil, err
	}
	return &App{
		Logger:  log,
		Wmi:     wmi.New(),
		Reg:     reg.New(),
		Files:   files.New(),
		IP:      ip.New(),
		Env:     env.New(),
		discord: dsRPC,
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

func (a *App) GetDiscordID() string {
	if a.discord != nil {
		r, err := a.discord.GetID()
		if err != nil {
			a.Logger.Errorf("error get discord id: %v", err)
			return "error"
		}
		return r
	}
	return "unknown"
}

func (a *App) GetDiscordUsername() string {
	if a.discord != nil {
		r, err := a.discord.GetUsername()
		if err != nil {
			a.Logger.Errorf("error get discord id: %v", err)
			return "error"
		}
		return r
	}
	return "unknown"
}

func (a *App) GetDiscordArray() string {
	u, id := a.GetDiscordUsername(), a.GetDiscordID()
	if u == "unknown" || id == "unknown" {
		return "unknown"
	}
	return fmt.Sprintf(`["%s","%s"]`, u, id)
}

func (a *App) ReadRegistry(category, path, key string) string {
	v, err := a.Reg.ReadReg(category, path, key)
	if err != nil {
		a.Logger.Errorf("error read registry: %v | %s | %s", err, path, key)
		return fmt.Sprintf("error read registry %s | %s", path, key)
	}
	return v
}

func (a *App) WriteRegistry(category, path, key, value string) string {
	err := a.Reg.WriteReg(category, path, key, value)
	if err != nil {
		a.Logger.Errorf("error write registry: %v | %s | %s | %s", err, path, key, value)
		return fmt.Sprintf("error write registry %s | %s | %s", path, key, value)
	}
	return "written"
}

func (a *App) DeleteRegistry(category, path, key string) string {
	err := a.Reg.DeleteReg(category, path, key)
	if err != nil {
		a.Logger.Errorf("error delete registry: %v | %s | %s", err, path, key)
		return fmt.Sprintf("error delete registry %s | %s", path, key)
	}
	return "deleted"
}

func (a *App) GetEnv(key string) string {
	return a.Env.Get(key)
}

func (a *App) SetEnv(key, value string) string {
	err := a.Env.Set(key, value)
	if err != nil {
		a.Logger.Errorf("error set env: %v | %s | %s", err, key, value)
		return fmt.Sprintf("error set env %s | %s", key, value)
	}
	return "setted"
}

func (a *App) WriteFile(path, data string) string {
	err := a.Files.WriteFile(path, data)
	if err != nil {
		a.Logger.Errorf("error write file: %v | %s | %s", err, path, data)
		return fmt.Sprintf("error write file %s | %s", path, data)
	}
	return "written"
}

func (a *App) ReadFile(path string) string {
	data, err := a.Files.ReadFile(path)
	if err != nil {
		a.Logger.Errorf("error read file: %v | %s", err, path)
		return fmt.Sprintf("error read file %s", path)
	}
	return data
}

func (a *App) DeleteFile(path string) string {
	err := a.Files.DeleteFile(path)
	if err != nil {
		a.Logger.Errorf("error delete file: %v | %s", err, path)
		return fmt.Sprintf("error delete file %s", path)
	}
	return "deleted"
}

func (a *App) GetIP() string {
	return a.IP.GetIp()
}

func (a *App) GetGeoIP() string {
	return a.IP.GetGeoIp()
}
