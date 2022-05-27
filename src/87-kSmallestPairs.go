package main

import (
	"container/heap"
	"fmt"
)

type maxHeap [][3]int

func (mh maxHeap) Len() int {
	return len(mh)
}

// 1. 大顶堆
//func (mh maxHeap) Less(i, j int) bool {
//	return mh[i][0] > mh[j][0]
//}

// 2. 小顶堆
func (mh maxHeap) Less(i, j int) bool {
	return mh[i][0] < mh[j][0]
}

func (mh maxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *maxHeap) Push(x interface{}) {
	*mh = append(*mh, x.([3]int))
}

func (mh *maxHeap) Pop() interface{} {
	old := *mh
	res := old[len(old)-1]
	*mh = old[:len(old)-1]
	return res
}

// 1. 大顶堆
//func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
//	mh := &maxHeap{}
//	for i := 0; i < len(nums1); i++ {
//		for j := 0; j < len(nums2); j++ {
//			sum := nums1[i] + nums2[j]
//			if mh.Len() < k || sum < (*mh)[0][0] {
//				heap.Push(mh, [3]int{sum, nums1[i], nums2[j]})
//			}
//			if mh.Len() > k {
//				heap.Pop(mh)
//			}
//		}
//	}
//	res := make([][]int, len(*mh))
//	for i := 0; i < len(*mh); i++ {
//		res[i] = []int{(*mh)[i][1], (*mh)[i][2]}
//	}
//	return res
//}

// 2. 小顶堆
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	mh := &maxHeap{}
	// 1. 初始堆
	for i := 0; i < k && i < len(nums1); i++ {
		sum := nums1[i] + nums2[0]
		heap.Push(mh, [3]int{sum, i, 0})
	}
	// 2. 对比，出堆
	res := [][]int{}
	cnt := 0
	for mh.Len() != 0 && cnt < k {
		temp := heap.Pop(mh).([3]int)
		res = append(res, []int{nums1[temp[1]], nums2[temp[2]]})
		if temp[2]+1 < len(nums2) {
			sum := nums1[temp[1]] + nums2[temp[2]+1]
			heap.Push(mh, [3]int{sum, temp[1], temp[2] + 1})
		}
		cnt++
	}
	return res
}

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	fmt.Println(kSmallestPairs(nums1, nums2, 3))
}
