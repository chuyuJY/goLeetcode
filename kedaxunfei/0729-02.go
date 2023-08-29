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
	nums1 := make([]int, n)
	nums2 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(sc, "%d", &nums1[i])
	}
	fmt.Fscanln(sc)
	for i := 0; i < n; i++ {
		fmt.Fscanf(sc, "%d", &nums2[i])
	}
	test2(nums1, nums2)
}

func test2(nums1, nums2 []int) {
	res := 0
	for i := 0; i < len(nums1); i++ {
		res += min(abs(nums1[i]-nums2[i]), abs(nums1[i]+nums2[i]))
	}
	fmt.Println(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
