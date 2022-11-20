package files

import (
	"fmt"
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

func (fl *Files) WriteFile(path, data string) error {
	if path[0:1] == "~" {
		home, _ := os.UserHomeDir()
		path = fmt.Sprint(home, path[1:])
	}

	spPath := strings.Split(path, "\\")
	fPath := strings.Join(spPath[:len(spPath)-1], "\\")

	if err := os.MkdirAll(fPath, os.ModeDir); err != nil {
		fl.logger.Errorf("MkDir Error: %s", err.Error())
		return fmt.Errorf("MkDir Error: %s", err.Error())
	}

	if _, err := os.Stat(path); err == nil {
		return nil
	}

	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		fl.logger.Errorf("Create Error: %s", err.Error())
		return fmt.Errorf("create Error: %s", err.Error())
	}

	_, err = f.Write([]byte(data))
	if err != nil {
		fl.logger.Errorf("Write Error: %s", err.Error())
		return fmt.Errorf("write Error: %s", err.Error())
	}

	nameptr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		_ = os.Remove(path)
		fl.logger.Errorf("Nameptr Error: %s", err.Error())
		return fmt.Errorf("nameptr Error: %s", err.Error())
	}

	err = syscall.SetFileAttributes(nameptr, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		_ = os.Remove(path)
		fl.logger.Errorf("Attribute Error: %s", err.Error())
		return fmt.Errorf("attribute Error: %s", err.Error())
	}
	return nil
}

func (fl *Files) ReadFile(path string) (string, error) {
	if path[0:1] == "~" {
		home, _ := os.UserHomeDir()
		path = fmt.Sprint(home, path[1:])
	}
	data, err := os.ReadFile(path)
	if err != nil {
		fl.logger.Errorf("Read Error: %s", err.Error())
		return "", err
	}
	return string(data[:]), nil
}

func (fl *Files) DeleteFile(path string) error {
	if path[0:1] == "~" {
		home, _ := os.UserHomeDir()
		path = fmt.Sprint(home, path[1:])
	}

	if err := os.RemoveAll(path); err != nil {
		fl.logger.Errorf("Delete Error: %s", err.Error())
		return err
	}
	return nil
}
