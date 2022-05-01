package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	index := head
	for i := 1; i < k; i++ {
		index = index.Next
	}
	for ; index.Next != nil; index = index.Next {
		head = head.Next
	}
	return head
}
