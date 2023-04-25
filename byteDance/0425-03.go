package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var mod int64 = 1e7 + 7

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	nums := []int{}
	for i := 0; i < n; i++ {
		var ai int
		fmt.Fscan(sc, &ai)
		nums = append(nums, ai)
	}
	fmt.Fscanln(sc)
	str, _ := sc.ReadString('\n')
	color := strings.TrimRight(str, "\r\n")
	test5(nums, color)
}

func test5(nums []int, color string) {
	redNums := []int{}
	blueNums := []int{}
	for i := 0; i < len(nums); i++ {
		if color[i] == 'R' {
			redNums = append(redNums, nums[i])
		} else {
			blueNums = append(blueNums, nums[i])
		}
	}
	sort.Ints(redNums)
	sort.Ints(blueNums)
	// 计算结果

}
