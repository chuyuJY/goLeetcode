package main

//type BSTIterator struct {
//	stack []*TreeNode
//}
//
//func Constructor(root *TreeNode) BSTIterator {
//	stack := []*TreeNode{}
//	for root != nil {
//		stack = append(stack, root)
//		root = root.Left
//	}
//	return BSTIterator{stack: stack}
//}
//
//func (this *BSTIterator) Next() int {
//	next := this.stack[len(this.stack)-1]
//	this.stack = this.stack[:len(this.stack)-1]
//	root := next.Right
//	for root != nil {
//		this.stack = append(this.stack, root)
//		root = root.Left
//	}
//	return next.Val
//}
//
//func (this *BSTIterator) HasNext() bool {
//	return len(this.stack) > 0
//}
