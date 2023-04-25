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
	nums := []int{}
	for i := 0; i < n; i++ {
		var ai int
		fmt.Fscan(sc, &ai)
		nums = append(nums, ai)
	}
	fmt.Fscanln(sc)
	test7(nums)
}

func test7(nums []int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		fmt.Println(0)
		return
	}
	target := sum / 2
	prefix := 0
	prefixs := []int{}
	for i := 0; i < len(nums); i++ {
		prefix += nums[i]
		prefixs = append(prefixs, prefix)
	}
	res := 0
	hashMap := map[int]int{}         // 记录已遍历过去的前缀和，非常巧妙啊
	for i := 0; i < len(nums); i++ { // 注意这里没有统计[0:i]是target的情况，因为会和[i:len(nums)-1]重复
		pre := prefixs[i] - target
		if val, exist := hashMap[pre]; exist {
			res += val
		}
		hashMap[prefixs[i]]++
	}
	fmt.Println(res)
}
