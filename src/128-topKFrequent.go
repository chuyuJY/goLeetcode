package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	hashMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]]++
	}
	mh := &minHeap{}
	for i, v := range hashMap {
		heap.Push(mh, []int{i, v})
		if mh.Len() > k {
			heap.Pop(mh)
		}
	}
	res := []int{}
	for _, num := range *mh {
		res = append(res, num[0])
	}
	return res
}

type minHeap [][]int

func (mh minHeap) Len() int {
	return len(mh)
}

func (mh minHeap) Less(i, j int) bool {
	return mh[i][1] < mh[j][1]
}

func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minHeap) Push(x interface{}) {
	*mh = append(*mh, x.([]int))
}

func (mh *minHeap) Pop() interface{} {
	old := *mh
	v := old[len(old)-1]
	*mh = old[:len(old)-1]
	return v
}
