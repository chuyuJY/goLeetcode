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
	rs := []int{}
	str, _ := sc.ReadString('\n')
	strs := strings.Split(strings.TrimRight(str, "\r\n"), " ")
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		rs = append(rs, num)
	}
	var cnt int
	fmt.Fscanln(sc, &cnt)
	test1(cnt, rs)
}

//func test1(cnt int, rs []int) {
//	sort.Ints(rs)
//	prefix := 0
//	prefixs := []int{}
//	for i := 0; i < len(rs); i++ {
//		prefix += rs[i]
//		prefixs = append(prefixs, prefix)
//	}
//	var isValid func(val int) bool
//	isValid = func(val int) bool {
//		index := sort.Search(len(rs), func(i int) bool {
//			return val < rs[i]
//		})
//		sum := 0
//		if index == 0 {
//			sum = len(rs) * val
//		} else {
//			sum = prefixs[index-1] + (len(rs)-index)*val
//		}
//		return sum <= cnt
//	}
//	left, right := 0, cnt+1
//	for left < right {
//		mid := (right-left)/2 + left
//		if isValid(mid) {
//			left = mid + 1
//		} else {
//			right = mid
//		}
//	}
//	if right == cnt+1 {
//		right = 0
//	}
//	fmt.Println(right - 1)
//}
