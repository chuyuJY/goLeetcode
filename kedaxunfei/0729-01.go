package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	vals := make([]int, n)
	str, _ := sc.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	strs := strings.Split(str, " ")
	for i := 0; i < n; i++ {
		vals[i], _ = strconv.Atoi(strs[i])
	}
	test1(vals)
}

func test1(vals []int) {
	res := 0
	index := 0
	for i := 0; i < len(vals); i++ {
		if vals[i] == -1 {
			index = i
			break
		}
	}
	cur := vals[0]
	for i := 0; i < index; i++ {
		cur = min(cur, vals[i])
	}
	if cur != -1 {
		res += cur
	}
	cur = vals[len(vals)-1]
	for i := index + 1; i < len(vals); i++ {
		cur = min(cur, vals[i])
	}
	if cur != -1 {
		res += cur
	}
	fmt.Println(res)
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
