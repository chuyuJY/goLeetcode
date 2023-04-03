package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%v\n", &n)
	sc := bufio.NewScanner(os.Stdin)
	nums := []int{}
	if sc.Scan() {
		str := sc.Text()
		strs := strings.Split(str, " ")
		for _, str := range strs {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}
	}
	fmt.Println(count(nums))
}

//func count(nums []int) int {
//	mod := 1000000007
//	sort.Ints(nums)
//	res := 0
//	var dfs func(index int, lastNum int)
//	dfs = func(index int, lastNum int) {
//		if index == len(nums) {
//			return
//		}
//		if nums[index]%lastNum == 0 {
//			res %= mod
//			res++
//			dfs(index+1, nums[index])
//		}
//		dfs(index+1, lastNum)
//	}
//	for i := 1; i < len(nums); i++ {
//		dfs(i, nums[i-1])
//	}
//	return res
//}
