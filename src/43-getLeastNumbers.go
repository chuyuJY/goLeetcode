package main

import "fmt"

func getLeastNumbers(arr []int, k int) []int {
	quickSort(arr, 0, len(arr)-1)
	//quickKthSort(arr, 0, len(arr)-1, k)
	return arr[:k]
}

// 快速排序
func quickSort(arr []int, left, right int) []int {
	// 终止条件

	if left >= right {
		// 注意此处不能是返回空，因为如果初始数组只有一个元素如：[16]，就会报错
		return arr
	}
	// 递归公式
	// 注意此处，不能是left+1，只有两个元素的情况，会无法对首元素进行排序
	i, j := left, right
	for i < j {
		// 注意此处的顺序，先判断j，在判断i ！！！ 为的是停止的时候arr[i] < arr[left]
		for i < j && arr[j] >= arr[left] {
			j--
		}
		for i < j && arr[i] <= arr[left] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	// 此时 i = j，且此时arr[i] < arr[left]
	arr[i], arr[left] = arr[left], arr[i]
	quickSort(arr, left, i-1)
	quickSort(arr, i+1, right)
	return arr
}

// 基于快速排序的数组划分（找到基准为k即可）
func quickKthSort(arr []int, left, right, k int) []int {
	if left >= right {
		return arr
	}

	i, j := left, right
	for i < j {
		for i < j && arr[j] > arr[left] {
			j--
		}
		for i < j && arr[i] < arr[left] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[left], arr[i] = arr[i], arr[left]
	// 和快速排序只有此处不同
	if i > k-1 {
		quickKthSort(arr, left, i-1, k)
	} else if i < k-1 {
		quickKthSort(arr, i+1, right, k)
	}
	return arr
}

func main() {
	nums := []int{9, 12, 56, 5, 24, 35, 5, 5, 6, 8, 43, 16}
	//nums = getLeastNumbers(nums, 5)
	nums = getLeastNumbers(nums, 5)
	fmt.Println(nums)
}
