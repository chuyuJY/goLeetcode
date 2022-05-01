package main

import "fmt"

func fastSort(nums []int, low, high int) []int {
	sort := func(nums []int, low, high int) int {
		t, p := low, nums[low]
		low++
		for {
			for ; nums[low] < p && low < high; low++ { // 最后一定是停在排序好的第一个比p大的元素上
			}
			for ; nums[high] >= p && high >= low; high-- { // 强烈注意此处的 high >= low ！！！要让high在从右往左第一个更小的元素上
			}
			if low >= high {
				break
			}
			nums[low], nums[high] = nums[high], nums[low]
		}
		nums[t], nums[high] = nums[high], nums[t]
		return high
	}
	if low < high {
		t := sort(nums, low, high)
		fastSort(nums, low, t-1)
		fastSort(nums, t+1, high)
	}
	return nums
}

func main() {
	nums := []int{9, 12, 56, 5, 24, 35, 5, 5, 6, 8, 43, 16}
	nums = fastSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
