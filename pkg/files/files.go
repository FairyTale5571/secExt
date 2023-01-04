package files

import (
	"fmt"
	"os"
	"strings"
	"syscall"
)

func WriteFile(path, data string) error {
	if path[0:1] == "~" {
		home, _ := os.UserHomeDir()
		path = fmt.Sprint(home, path[1:])
	}

	spPath := strings.Split(path, "\\")
	fPath := strings.Join(spPath[:len(spPath)-1], "\\")

	if err := os.MkdirAll(fPath, os.ModeDir); err != nil {
		return fmt.Errorf("MkDir Error: %s", err.Error())
	}

	if _, err := os.Stat(path); err == nil {
		return nil
	}

	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("create Error: %s", err.Error())
	}

	_, err = f.Write([]byte(data))
	if err != nil {
		return fmt.Errorf("write Error: %s", err.Error())
	}

	nameptr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		_ = os.Remove(path)
		return fmt.Errorf("nameptr Error: %s", err.Error())
	}

	err = syscall.SetFileAttributes(nameptr, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		_ = os.Remove(path)
		return fmt.Errorf("attribute Error: %s", err.Error())
	}
	return nil
}

func ReadFile(path string) (string, error) {
	if path[0:1] == "~" {
		home, _ := os.UserHomeDir()
		path = fmt.Sprint(home, path[1:])
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data[:]), nil
}

func DeleteFile(path string) error {
	if path[0:1] == "~" {
		home, _ := os.UserHomeDir()
		path = fmt.Sprint(home, path[1:])
	}

	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}
