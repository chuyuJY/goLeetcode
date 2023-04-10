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
	var n, k int
	fmt.Fscanln(sc, &n, &k)
	nums := []int{}
	str, _ := sc.ReadString('\n')
	strs := strings.Split(strings.TrimRight(str, "\r\n"), " ")
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	fmt.Println(test(n, k, nums))
}

func test(n, k int, nums []int) int {
	res := 0
	hashMap := map[int]int{}
	count := 0
	left, right := 0, 0
	for right < n {
		for left < right && count > k {
			hashMap[nums[left]]--
			if hashMap[nums[left]] == 0 {
				count--
			}
			left++
		}
		res = max(res, right-left)
		if val, exist := hashMap[nums[right]]; !exist || val == 0 {
			count++
		}
		hashMap[nums[right]]++
		right++
	}
	if count <= k {
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
