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
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(sc, "%d", &nums[i])
	}
	fmt.Fscanln(sc)
	var s int
	fmt.Fscanln(sc, &s)
	test3(nums, s)
}

func test3(nums []int, s int) {
	res := s - 1
	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}
	for _, num := range nums {
		if num < s {
			if hashMap[s-num] {
				res -= 1
			} else {
				res -= 2
			}
		}
	}
	fmt.Println(res)
}
