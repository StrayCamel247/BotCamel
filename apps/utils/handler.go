package utils

import (
	"strings"
)

func init() {

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
