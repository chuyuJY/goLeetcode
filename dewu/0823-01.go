package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	hs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(sc, &hs[i])
	}
	fmt.Fscanln(sc)
	test1(hs)
}

func test1(hs []int) {
	sort.Ints(hs)
	res := hs[1] - hs[0]
	cur := hs[len(hs)-1]
	left := len(hs) - 2
	for left >= 0 {
		res = max(res, cur-hs[left])
		cur = hs[left]
		left -= 2
	}
	cur = hs[len(hs)-1]
	right := len(hs) - 3
	for right >= 0 {
		res = max(res, cur-hs[right])
		cur = hs[right]
		right -= 2
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
