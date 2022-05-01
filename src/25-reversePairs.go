package main

import "fmt"

var temp []int

func reversePairs(nums []int) int {
	length := len(nums)
	temp = make([]int, length)
	return mergeSort(nums, 0, length-1)
}
func mergeSort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	// 划分
	// (left+right)/2这样写是有可能发生整型溢出的，因此最好是left+(right-left)/2
	//mid := (left + right) / 2
	mid := left + (right-left)/2
	i, j := left, mid+1
	res := mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right)
	// 归并
	for cnt := left; cnt <= right; cnt++ {
		// 一次循环只排序一个元素
		if i == mid+1 { // i 完毕
			temp[cnt] = nums[j]
			j++
		} else if j == right+1 || nums[i] <= nums[j] { // j完毕或nums[i]<=nums[j]
			temp[cnt] = nums[i]
			i++
		} else {
			temp[cnt] = nums[j]
			j++
			res += mid - i + 1
		}
	}
	// 将排序完毕的数组给原数组
	for k := left; k <= right; k++ {
		nums[k] = temp[k]
	}
	return res
}

func main() {
	nums := []int{7, 5, 8, 6, 4, 3}
	fmt.Println(reversePairs(nums))
}
