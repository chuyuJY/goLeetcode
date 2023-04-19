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
	aNums, bNums := []int{}, []int{}
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(sc, &num)
		aNums = append(aNums, num)
	}
	fmt.Fscanln(sc)
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(sc, &num)
		bNums = append(bNums, num)
	}
	fmt.Fscanln(sc)
	test3(n, aNums, bNums)
}

func test3(n int, aNums, bNums []int) {
	v := [3][]int{} // 按 b 数组 0、1、2 分组
	for i := 0; i < n; i++ {
		v[bNums[i]] = append(v[bNums[i]], aNums[i])
	}
	res := 0
	c := 0
	for i := 0; i < 3; i++ {
		sort.Ints(v[i])
		for j := 0; j < len(v[i]); j++ {
			c++
			res += abs(v[i][j], c)
		}
	}
	fmt.Println(res)
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
