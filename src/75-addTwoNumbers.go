package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Stack, l2Stack := []int{}, []int{}
	var head *ListNode
	for l1 != nil {
		l1Stack = append(l1Stack, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		l2Stack = append(l2Stack, l2.Val)
		l2 = l2.Next
	}
	in := 0
	for i := 0; i < len(l1Stack) || i < len(l2Stack) || in != 0; i++ {
		if i < len(l1Stack) {
			in += l1Stack[len(l1Stack)-i-1]
		}
		if i < len(l2Stack) {
			in += l2Stack[len(l2Stack)-i-1]
		}
		tempHead := &ListNode{
			in % 10,
			head,
		}
		in /= 10
		head = tempHead
	}
	return head
}
func main() {
	node := &ListNode{
		0,
		nil,
	}
	var head *ListNode
	//head = node
	//node = head
	//head.Val = node.Val
	fmt.Println(&node, &head)
}
