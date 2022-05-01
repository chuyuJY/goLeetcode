package main

import (
	"fmt"
	"strconv"
)

// 快速排序
func minNumber(nums []int) string {
	cnt := len(nums)
	str := make([]string, cnt)
	for i := 0; i < cnt; i++ {
		str[i] = strconv.Itoa(nums[i])
	}
	str = selfQuickSort(str, 0, cnt-1)
	var res string
	for i := 0; i < cnt; i++ {
		res += str[i]
	}
	return res
}

func selfQuickSort(str []string, left, right int) []string {
	if left >= right {
		// 注意此处不能是返回空，因为如果初始数组只有一个元素如：[16]，就会报错
		return str
	}
	i, j := left, right
	for i < j {
		// 注意此处一定要先j，不然就不能交换i和left了，因为i==j会发生在>temp的位置上
		// 相等的就不管了，因为左右都会继续递归排序
		for i < j && str[j]+str[left] >= str[left]+str[j] {
			j--
		}
		for i < j && str[i]+str[left] <= str[left]+str[i] {
			i++
		}
		str[i], str[j] = str[j], str[i]
	}
	str[left], str[i] = str[i], str[left]
	selfQuickSort(str, left, i-1)
	selfQuickSort(str, i+1, right)
	return str
}

func main() {
	str := []int{10, 2}
	fmt.Println(minNumber(str))
}
