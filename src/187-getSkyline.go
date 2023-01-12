package main

import (
	"container/heap"
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	res := [][]int{}
	// 扫描线
	line := [][2]int{}
	for _, building := range buildings {
		line = append(line, [2]int{building[0], -building[2]}) // 左端点
		line = append(line, [2]int{building[1], building[2]})  // 右端点
	}
	sort.Slice(line, func(i, j int) bool {
		if line[i][0] == line[j][0] {
			return line[i][1] < line[j][1] // 左端点: 由高到低; 右端点: 由低到高
		}
		return line[i][0] < line[j][0] // 由左至右
	})
	// 大顶堆+遍历
	mh := &maxHeap{}
	prev, cur := 0, 0
	heap.Push(mh, cur)
	for _, point := range line {
		x, y := point[0], point[1]
		if y < 0 {
			heap.Push(mh, -y)
		} else {
			heap.Remove(mh, mh.Find(y))
		}

		if mh.Len() > 0 {
			cur = (*mh)[0]
		}
		if cur != prev {
			res = append(res, []int{x, cur})
			prev = cur
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

func (mh *maxHeap) Find(v interface{}) int {
	idx := -1
	for i := 0; i < len(*mh); i++ {
		if (*mh)[i] == v.(int) {
			idx = i
			break
		}
	}
	return idx
}
