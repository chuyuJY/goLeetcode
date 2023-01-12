package main

import "container/heap"

type MedianFinder struct {
	minH *minHeap
	maxH *maxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minH: &minHeap{},
		maxH: &maxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minH.Len() == 0 {
		heap.Push(this.minH, num)
	} else if this.maxH.Len() == 0 {
		bottom := heap.Pop(this.minH).(int)
		if num > bottom {
			num, bottom = bottom, num
		}
		heap.Push(this.minH, bottom)
		heap.Push(this.maxH, num)
	} else {
		if this.minH.Len() > this.maxH.Len() {
			bottom := heap.Pop(this.minH).(int)
			if bottom < num {
				num, bottom = bottom, num
			}
			heap.Push(this.minH, bottom)
			heap.Push(this.maxH, num)
		} else {
			top := heap.Pop(this.maxH).(int)
			if top > num {
				top, num = num, top
			}
			heap.Push(this.minH, num)
			heap.Push(this.maxH, top)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minH.Len() == this.maxH.Len() {
		top, bottom := heap.Pop(this.maxH).(int), heap.Pop(this.minH).(int)
		heap.Push(this.maxH, top)
		heap.Push(this.minH, bottom)
		return float64(top+bottom) / 2.0
	} else {
		bottom := heap.Pop(this.minH).(int)
		heap.Push(this.minH, bottom)
		return float64(bottom)
	}
}

// type minHeap []int
// type maxHeap []int

func (mh minHeap) Len() int {
	return len(mh)
}

func (mh minHeap) Less(i, j int) bool {
	return mh[i] < mh[j]
}

func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minHeap) Push(v interface{}) {
	*mh = append(*mh, v.(int))
}

func (mh *minHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	v := old[n-1]
	*mh = old[:n-1]
	return v
}

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
