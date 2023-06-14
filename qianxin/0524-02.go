package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}

	dummy := &ListNode{0, head}
	first, second := dummy, dummy

	for i := 0; i < n; i++ {
		if first.Next == nil {
			return head
		}
		first = first.Next
	}

	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printList(head)

	n := 2
	newHead := removeNthFromEnd(head, n)
	printList(newHead)
}
