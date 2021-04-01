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
	Remark  string
	Command string
}

// 获取指令匹配项
func (r *Info) key() []string {
	return append(r.Keys, r.Command)
}
