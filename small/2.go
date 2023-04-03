package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var t int
	fmt.Scanf("%v\n", &t)
	for i := 0; i < t; i++ {
		var n, k int
		fmt.Scanf("%v %v\n", &n, &k)
		nums := []int{}
		sc := bufio.NewScanner(os.Stdin)
		if sc.Scan() {
			str := sc.Text()
			strs := strings.Split(str, " ")
			for _, str := range strs {
				num, _ := strconv.Atoi(str)
				nums = append(nums, num)
			}
		}
		fmt.Println(test2(nums, k))
	}
}

func test2(nums []int, k int) []int {
	if len(nums) < 1 {
		return nums
	}
	if len(nums) < k {
		k = len(nums)
	}
	heap := nums[:k]
	for i := (k - 2) / 2; i >= 0; i-- {
		sort1(heap, i, k)
	}
	for i := k; i < len(nums); i++ {
		nums[i-k] = heap[0]
		heap[0] = nums[i]
		sort1(heap, 0, k)
	}

	for i := 0; i < k; i++ {
		nums[len(nums)-k+i] = heap[0]
		heap[0] = heap[k-1-i]
		sort1(heap, 0, k-1-i)
	}
	return nums
}

func sort1(nums []int, root int, length int) {
	for {
		child := root*2 + 1
		if child >= length {
			break
		}
		if child+1 < length && nums[child+1] < nums[child] {
			child++
		}
		if nums[root] <= nums[child] {
			break
		}
		nums[root], nums[child] = nums[child], nums[root]
		root = child
	}
}
