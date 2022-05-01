package main

//import "container/list"
//
//type CQueue struct {
//	stack1, stack2 *list.List // list包的List函数，是一个结构体，双向链表，充当栈
//}
//
//func Constructor() CQueue {
//	return CQueue{
//		stack1: list.New(), // 创建空链表 // 用stack1存数据，把删除的数据放到stack2中
//		stack2: list.New(),
//	}
//}
//
//func (this *CQueue) AppendTail(value int) {
//	this.stack1.PushBack(value)
//}
//
//func (this *CQueue) DeleteHead() int {
//	if this.stack2.Len() == 0 {
//		for this.stack1.Len() != 0 { // 若stack2已经没有元素，则需要把stack1中的全都倒入
//			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back())) // Remove会返回节点的值
//		}
//	}
//	if this.stack2.Len() != 0 { // 若stack2不为空，则直接弹出并删除即可
//		return this.stack2.Remove(this.stack2.Back()).(int) //注意V大写，并使用类型断言转为int型
//	}
//	return -1 // 此时1和2都没有元素了，返回默认值即可
//}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
