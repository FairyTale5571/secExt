package helpers

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func ConvertSize(size uint64) (uint64, string) {
	if size < 1000*1024 {
		return size / 1024, "KB"
	}
	if size < 1000*1048576 {
		return size / 1048576, "MB"
	}
	if size < 1000*1073741824 {
		return size / 1073741824, "GB"
	} else {
		return size / 1099511627776, "TB"
	}
}

func Struct2JSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func IsAdmin() string {
	file, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	defer file.Close()
	if err != nil {
		return "false"
	}
	return "true"
}

func EnsureDir(fileName string) error {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return merr
		}
	}
	return nil
}
