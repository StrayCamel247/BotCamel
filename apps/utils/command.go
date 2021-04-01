package utils

import (
	"fmt"
	"strconv"
)

var commandNum int64

func init() {

}

// 项目启动时-分发递增得command
func FetchCommandNum() string {
	defer func() { commandNum++ }()
	// 串行分发序号-转为16进制
	return fmt.Sprintf("0x%s", strconv.FormatInt(commandNum, 16))
}

// 指令信息
type Info struct {
	Keys    []string
	Remark  string `json:"remark"`
	Command string `json:"com"`
}

func RemakrFormat(s string) string {
	return fmt.Sprintf("	%s	", s)
}

// 获取指令匹配项
func (r *Info) Key() []string {
	return append(r.Keys, r.Command)
}
