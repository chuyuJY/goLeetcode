package main

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

var str = []string{"0", "6", "8", "9"}
var passwd string = "0000069899668900"

// 1. 回溯法
func dfs(hack string, k int) bool {
	// 到达叶节点
	if k == 0 {
		if hack == passwd {
			// 用绿色标记找出的passwd
			color.Green(hack)
			return true
		}
		fmt.Println(hack)
		return false
	} else {
		flag := false
		for i := 0; i < len(str) && flag == false; i++ {
			// 尝试新的密码
			newHack := strings.Join([]string{hack, str[i]}, "")
			// 已找到passwd之后，避免不必要递归，找到passwd即停止
			flag = flag || dfs(newHack, k-1)
		}
		return flag
	}
}

func main() {
	fmt.Println(dfs("", len(passwd)))
}
