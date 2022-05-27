//package main
//
//import "container/heap"
//
//type maxHeap []int // 大顶堆
//type minHeap []int // 小顶堆
//// 返回堆大小
//func (m maxHeap) Len() int {
//	return len(m)
//}
//func (m minHeap) Len() int {
//	return len(m)
//}
//
//// 决定大优先还是小优先
//func (m maxHeap) Less(i, j int) bool {
//	return m[i] > m[j]
//}
//func (m minHeap) Less(i, j int) bool {
//	return m[i] < m[j]
//}
//
//// 交换
//func (m maxHeap) Swap(i, j int) {
//	m[i], m[j] = m[j], m[i]
//}
//func (m minHeap) Swap(i, j int) {
//	m[i], m[j] = m[j], m[i]
//}
//
//// Push 在堆尾添加一个元素
//func (m *maxHeap) Push(v interface{}) {
//	*m = append(*m, v.(int))
//}
//func (m *minHeap) Push(v interface{}) {
//	*m = append(*m, v.(int))
//}
//
//// Pop 删除末尾元素
//func (m *maxHeap) Pop() interface{} {
//	old := *m
//	n := len(old)
//	v := old[n-1]
//	*m = old[:n-1]
//	return v
//}
//func (m *minHeap) Pop() interface{} {
//	old := *m
//	n := len(old)
//	v := old[n-1]
//	*m = old[:n-1]
//	return v
//}
//
//// 堆
//type MedianFinder struct {
//	maxH *maxHeap
//	minH *minHeap
//}
//
///** initialize your data structure here. */
////func Constructor() MedianFinder {
////	return MedianFinder{
////		maxH: new(maxHeap),
////		minH: new(minHeap),
////	}
////}
//
//// AddNum 插入元素
////
//func (this *MedianFinder) AddNum(num int) {
//	// 总体两种情况：
//	// 1.两个堆长度相等，此时要添加到小顶堆
//	// 2.小顶堆长度大于大顶堆，这时候要添加到大顶堆
//	if this.maxH.Len() == this.minH.Len() {
//		if this.minH.Len() == 0 || num > (*this.minH)[0] {
//			heap.Push(this.minH, num)
//		} else {
//			heap.Push(this.maxH, num)
//			// 得到大顶堆的最大值
//			top := heap.Pop(this.maxH).(int)
//			heap.Push(this.minH, top)
//		}
//	} else {
//		if num > (*this.minH)[0] {
//			heap.Push(this.minH, num)
//			bottle := heap.Pop(this.minH).(int)
//			heap.Push(this.maxH, bottle)
//		} else {
//			heap.Push(this.maxH, num)
//		}
//	}
//}
//
//func (this *MedianFinder) FindMedian() float64 {
//	if this.minH.Len() == this.maxH.Len() {
//		return float64((*this.minH)[0])/2.0 + float64((*this.maxH)[0])/2.0
//	} else {
//		return float64((*this.minH)[0])
//	}
//}
