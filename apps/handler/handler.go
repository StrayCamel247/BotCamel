package handler

import (
	"fmt"
	"os"
	"strings"
)

func int() {

}

// 字符按-数组匹配
func EqualFolds(s string, sArray []string) bool {
	for _, v := range sArray {
		if strings.EqualFold(s, v) {
			return true
		}

	}
	return false
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
