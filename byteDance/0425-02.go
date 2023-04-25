package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscanln(sc, &n, &k)
	nums := []int{}
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(sc, &num)
		nums = append(nums, num)
	}
	fmt.Fscanln(sc)
	test4(nums, k)
}

func test4(nums []int, k int) {
	prefix := 0
	prefixs := []int{0}
	for _, num := range nums {
		prefix += num
		prefixs = append(prefixs, prefix)
	}
	res := 0
	for i := k; i < len(prefixs); i++ {
		res = max(res, prefixs[i]-prefixs[i-k])
	}

	// 重新计算删除某个数字
	for i := k + 1; i < len(prefixs); i++ {
		for j := i; j >= i-k; j-- {
			res = max(res, prefixs[i]-nums[j-1]-prefixs[i-k])
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func test4(nums []int, k int) {
//	res := 0
//	for i := 0; i < len(nums); i++ {
//		temp := []int{}
//		temp = append(temp, nums[0:i]...)
//		temp = append(temp, nums[i+1:]...)
//		curRes := calRes(temp, k)
//		if curRes > res {
//			res = curRes
//		}
//	}
//	fmt.Println(res)
//}

//func calRes(nums []int, k int) int {
//	prefix := 0
//	prefixs := []int{0}
//	for _, num := range nums {
//		prefix += num
//		prefixs = append(prefixs, prefix)
//	}
//	res := 0
//	for i := k; i < len(prefixs); i++ {
//		if prefixs[i]-prefixs[i-k] > res {
//			res = prefixs[i] - prefixs[i-k]
//		}
//	}
//	return res
//}
