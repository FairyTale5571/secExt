package helpers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
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

func GenerateGUID() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		// handle error
	}
	uuid = fmt.Sprintf("%x%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:13], b[13:])
	return
}

func IsAdmin() string {
	file, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	defer file.Close()
	if err != nil {
		return "false"
	}
	return "true"
}
