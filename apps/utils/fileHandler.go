package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// ReadFile 读取文件
// 读取失败返回 nil
func ReadFile(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.WithError(err).WithField("util", "ReadFile").Errorf("unable to read '%s'", path)
		return nil
	}
	return bytes
}

// 检查文件是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("File reading error", err)
	return false
}
