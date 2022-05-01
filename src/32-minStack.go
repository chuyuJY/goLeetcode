package main

// 和08-CQueue类似，可以一起看
//type MinStack struct {
//	stack, minStack *list.List
//}
//
///** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{
//		stack:    list.New(),
//		minStack: list.New(),
//	}
//}
//
//func (this *MinStack) Push(x int) {
//	this.stack.PushBack(x)
//	// 保证minStack中的元素为非严格降序
//	if this.minStack.Len() == 0 || this.minStack.Back().Value.(int) >= x {
//		this.minStack.PushBack(x)
//	}
//}
//
//func (this *MinStack) Pop() {
//	y := this.stack.Remove(this.stack.Back())
//	if y == this.minStack.Back().Value {
//		this.minStack.Remove(this.minStack.Back())
//	}
//}
//
//func (this *MinStack) Top() int {
//	return this.stack.Back().Value.(int)
//}
//
//func (this *MinStack) Min() int {
//	return this.minStack.Back().Value.(int)
//}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

type MinStack struct {
	stack, minStack []int
}

//func Constructor() MinStack {
//	return MinStack{
//		stack:    []int{},
//		minStack: []int{},
//	}
//}
//func (this *MinStack) Push(x int) {
//	this.stack = append(this.stack, x)
//	if len(this.minStack) == 0 || this.minStack[len(this.minStack)-1] >= x {
//		this.minStack = append(this.minStack, x)
//	}
//}
//func (this *MinStack) Pop() {
//	length := len(this.stack)
//	y := this.stack[length-1]
//	// 切片是 左闭右开，因此是length-1
//	this.stack = this.stack[:length-1]
//	if y == this.minStack[len(this.minStack)-1] {
//		this.minStack = this.minStack[:len(this.minStack)-1]
//	}
//}
//func (this *MinStack) Top() int {
//	return this.stack[len(this.stack)-1]
//}
//func (this *MinStack) Min() int {
//	return this.minStack[len(this.minStack)-1]
//}
