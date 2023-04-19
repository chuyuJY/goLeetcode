package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 杯数
	var n int
	fmt.Scanf("%v\n", &n)
	sc := bufio.NewReader(os.Stdin)
	// 容量
	caps := []int{}
	str, _ := sc.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	strs := strings.Split(str, " ")
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		caps = append(caps, num)
	}
	// 初始水量
	begins := []int{}
	str, _ = sc.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	strs = strings.Split(str, " ")
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		begins = append(begins, num)
	}
	// 法力值/ml
	prices := []int{}
	str, _ = sc.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	strs = strings.Split(str, " ")
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		prices = append(prices, num)
	}
	// 练习数
	//var m int
	//fmt.Scanf("%v\n", &m)	// 本地不会报错，但是系统内会报错，很。。。睿智
	str, _ = sc.ReadString('\n')

	//  倒满哪个杯子
	str, _ = sc.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	strs = strings.Split(str, " ")
	for _, str := range strs {
		target, _ := strconv.Atoi(str)
		fmt.Printf("%v ", test10(target, caps, begins, prices))
	}
}

func test10(target int, caps []int, begins []int, prices []int) int {
	// 保持原样, 创建副本
	tmp := make([]int, len(begins))
	copy(tmp, begins)
	suffix := 0
	for i := target - 1; i >= 0; i-- {
		suffix += caps[i] - begins[i]
	}
	res := (caps[target-1] - begins[target-1]) * prices[target-1]
	for i := 0; i < target; i++ {
		res = min(res, suffix*prices[i])
		suffix -= caps[i] - begins[i]
	}
	begins = tmp
	return res
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
