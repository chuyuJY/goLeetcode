package main

//// 自定义堆
//type PriorityQueue []int
//
//// 以下是heap.Interface的模板方法，需要熟练掌握
//func (p PriorityQueue) Len() int {
//	return len(p)
//}
//
//// 比较 (i < j)
//func (p PriorityQueue) Less(i, j int) bool {
//	return p[i] < p[j]
//}
//
//// 交换 (i < j)
//func (p PriorityQueue) Swap(i, j int) {
//	p[i], p[j] = p[j], p[i]
//}
//
//// 添加
//func (p *PriorityQueue) Push(x interface{}) {
//	*p = append(*p, x.(int))
//}
//
//// 删除
//func (p *PriorityQueue) Pop() interface{} {
//	oldHeap := *p
//	n := len(oldHeap)
//	tail := oldHeap[n-1]
//	*p = oldHeap[:n-1]
//	return tail
//}
//
//// 正式题解
//type KthLargest struct {
//	MinHeap *PriorityQueue
//	Size    int
//}
//
//func Constructor(k int, nums []int) KthLargest {
//	minHeap := &PriorityQueue{}
//	kth := KthLargest{
//		MinHeap: minHeap,
//		Size:    k,
//	}
//	// 构造初始堆
//	for _, num := range nums {
//		kth.Add(num)
//	}
//	return kth
//}
//
//// Add 往小顶堆添加数据，并返回topK
//func (this *KthLargest) Add(val int) int {
//	// 若堆中元素不满，或者新元素大于堆顶元素，可以直接入堆
//	if this.Size > this.MinHeap.Len() || val > (*this.MinHeap)[0] {
//		heap.Push(this.MinHeap, val)
//	}
//	if this.Size < this.MinHeap.Len() {
//		heap.Pop(this.MinHeap)
//	}
//	return (*this.MinHeap)[0]
//}
