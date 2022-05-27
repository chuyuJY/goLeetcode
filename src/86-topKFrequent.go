package main

import (
	"container/heap"
	"fmt"
)

// []int{数字, 个数}
type freHeap [][]int

func (fp freHeap) Len() int {
	return len(fp)
}

func (fp freHeap) Less(i, j int) bool {
	// 小顶堆，按照频率
	return fp[i][1] < fp[j][1]
}

func (fp freHeap) Swap(i, j int) {
	fp[i], fp[j] = fp[j], fp[i]
}

func (fp *freHeap) Push(x interface{}) {
	*fp = append(*fp, x.([]int))
}

func (fp *freHeap) Pop() interface{} {
	old := *fp
	v := old[len(old)-1]
	*fp = old[:len(old)-1]
	return v
}

func topKFrequent(nums []int, k int) []int {
	hashMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]]++
	}
	fp := &freHeap{}
	for key, val := range hashMap {
		heap.Push(fp, []int{key, val})
		if fp.Len() > k {
			// 不断将最小的弹出堆
			heap.Pop(fp)
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		// 末尾的k个即为频率最高的元素
		// 注意此处是 heap
		res[i] = heap.Pop(fp).([]int)[0]
	}
	return res
}

func main() {
	array := []int{3, 0, 1, 0}
	fmt.Println(topKFrequent(array, 3))
}
