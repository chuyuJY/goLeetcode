package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	nums := []int64{}
	for i := 0; i < n; i++ {
		var num int64
		fmt.Fscan(sc, &num)
		nums = append(nums, num)
	}
	test4(nums)
}

func test4(nums []int64) {
	t := []int{}
	index := 0
	for index < len(nums) {
		count := 0
		if nums[index] == 1 {
			for index < len(nums) && nums[index] == 1 {
				count++
				index++
			}
		} else {
			index++
		}
		t = append(t, count) // 0: 表示该位置是 >1 的数字, >0 : 表示该位置是连续 1 的个数
	}
	res := 0
	for i := 0; i < len(t); i++ {
		temp := 0
		if t[i] == 0 {
			pre, suf := 0, 0 // 统计前后的 1 的个数
			if i > 0 {
				pre = t[i-1]
			}
			if i+1 < len(t) {
				suf = t[i+1]
			}
			temp = (pre/2+1)*(suf/2+1) + (pre+1)/2*(suf+1)/2 // 左边的偶数*右边的偶数 + 左边的奇数*右边的奇数
		} else {
			if t[i]&1 == 0 {
				temp = (t[i] + 2) / 2 * t[i] / 2
			} else {
				temp = (t[i] + 1) / 2 * (t[i] + 1) / 2
			}
		}
		res += temp
	}
	fmt.Println(res)
}
