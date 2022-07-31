package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := []string{}
	combination := []string{}
	var backTracking func(s string)
	backTracking = func(s string) {
		if s == "" && len(combination) == 4 {
			res = append(res, strings.Join(combination, "."))
			return
		}
		// 剪枝
		if len(combination) >= 4 {
			return
		}
		for i := 1; i <= len(s); i++ {
			if isIpNum(s[:i]) {
				combination = append(combination, s[:i])
				backTracking(s[i:])
				combination = combination[:len(combination)-1]
			}
		}
	}
	backTracking(s)
	return res
}

func isIpNum(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	ipNum, _ := strconv.Atoi(s)
	// 可排除四位及以上的可能
	if ipNum > 255 {
		return false
	}
	return true
}

func main() {
	str := ""
	fmt.Println(isIpNum(str))
}
