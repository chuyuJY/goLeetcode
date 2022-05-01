package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func deleteNode(head *ListNode, val int) *ListNode {
	index := head
	if index.Val == val {
		return head.Next
	}
	for index.Next.Val != val {
		index = index.Next
	}
	index.Next = index.Next.Next
	return head
}
