package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"github.com/fairytale5571/secExt/pkg/logger"
)

type Files struct {
	logger *logger.Wrapper
}

func New() *Files {
	return &Files{
		logger: logger.New("files"),
	}
}

func (fl *Files) WriteFile(path, data string) string {
	if path[0:1] == "~" {
		home, _ := os.UserHomeDir()
		path = fmt.Sprint(home, path[1:])
	}

	spPath := strings.Split(path, "\\")
	fPath := strings.Join(spPath[:len(spPath)-1], "\\")
	err := os.MkdirAll(fPath, os.ModeDir)
	if err != nil {
		return fmt.Sprintf("MkDir Error: %s", err.Error())
	}

	if _, err := os.Stat(path); err == nil {
		return "Already exist"
	}

	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return fmt.Sprintf("Create Error: %s", err.Error())
	}

	_, err = f.Write([]byte(data))
	if err != nil {
		return fmt.Sprintf("Write Error: %s", err.Error())
	}

	nameptr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		_ = os.Remove(path)
		return fmt.Sprintf("Nameptr Error: %s", err.Error())
	}

	err = syscall.SetFileAttributes(nameptr, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		_ = os.Remove(path)
		return fmt.Sprintf("Attribute Error: %s", err.Error())
	}
	return "Written"
}

func (fl *Files) ReadFile(path string) string {
	if path[0:1] == "~" {
		home, _ := os.UserHomeDir()
		path = fmt.Sprint(home, path[1:])
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err.Error()
	}
	return string(data[:])
}

func (fl *Files) DelFile(path string) string {
	if path[0:1] == "~" {
		home, _ := os.UserHomeDir()
		path = fmt.Sprint(home, path[1:])
	}

	resp := "Deleted"
	if err := os.RemoveAll(path); err != nil {
		resp = err.Error()
	}
	return resp
}
