package main

import (
	"container/heap"
	"math"
)

func kClosest(points [][]int, k int) [][]int {
	mh := &maxHeap{}
	for _, point := range points {
		if mh.Len() < k || calDistance(point) < calDistance((*mh)[0]) {
			heap.Push(mh, point)
		}
		if mh.Len() > k {
			heap.Pop(mh)
		}
	}
	return *mh
}

func calDistance(point []int) float64 {
	return math.Sqrt(float64(point[0]*point[0] + point[1]*point[1]))
}

// type maxHeap [][]int

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return calDistance(m[i]) > calDistance(m[j])
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(v interface{}) {
	*m = append(*m, v.([]int))
}

func (m *maxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	v := old[n-1]
	*m = old[:n-1]
	return v
}
