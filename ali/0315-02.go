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
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(sc, &a)
		nums[i] = a
	}
	fmt.Fscanln(sc)
	test4(nums)
}

func test4(nums []int) {
	var res int64
	sort.Ints(nums)
	left, right := 0, 0
	for right < len(nums) {
		for left <= right && nums[right]-nums[left] > 1 {
			left++
		}
		if right-left >= 2 && nums[right]-nums[left] == 1 {
			// 统计所有以right结尾的合法三元组
			cur := left
			for right-cur >= 2 && nums[right]-nums[cur] == 1 {
				res += int64(right - cur - 1)
				cur++
			}
		}
		right++
	}
	fmt.Println(res)
}
