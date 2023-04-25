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
	a := []int{}
	for i := 0; i < n; i++ {
		var ai int
		fmt.Fscan(sc, &ai)
		a = append(a, ai)
	}
	fmt.Fscanln(sc)
	test4(a)
}

func test4(a []int) {
	sort.Ints(a)
	res := 0
	left, right := 0, 0
	for right < len(a) {
		for right < len(a) && a[right] == a[left] {
			right++
		}
		if a[left]&1 == 1 && (right-left)&1 == 1 {
			res += right - left
		}
		left = right
	}
	fmt.Println(res)
}
