package handler

import "strings"

func int() {

}
func EqualFolds(s string, sArray []string) bool {
	for _, v := range sArray {
		if strings.EqualFold(s, v) {
			return true
		}

	}
	return false
}
