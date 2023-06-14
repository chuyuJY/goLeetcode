package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(maxTime([]int{100, 100, 100, 100, 95, 10}))
}

func maxTime(batteries []int) int {
	res := 0
	if len(batteries) < 5 {
		return res
	}
	mh := &maxHeap{}
	for i := 0; i < len(batteries); i++ {
		heap.Push(mh, batteries[i])
	}
	for mh.Len() >= 5 {
		temp := []int{}
		for i := 0; i < 5; i++ {
			temp = append(temp, heap.Pop(mh).(int))
		}
		for i := 0; i < len(temp); i++ {
			temp[i]-- // 真坑啊。。。
		}
		res++
		for i := 0; i < len(temp); i++ {
			if temp[i] > 0 {
				heap.Push(mh, temp[i])
			}
		}
	}
	return res
}

type maxHeap []int

func (mh maxHeap) Len() int {
	return len(mh)
}

func (mh maxHeap) Less(i, j int) bool {
	return mh[i] > mh[j]
}

func (mh maxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *maxHeap) Push(v interface{}) {
	*mh = append(*mh, v.(int))
}

func (mh *maxHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	v := old[n-1]
	*mh = old[:n-1]
	return v
}
