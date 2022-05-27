package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	second := splitList(head)
	return mergeList(sortList(head), sortList(second))
}

func splitList(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	second := slow.Next
	slow.Next = nil
	return second
}

func mergeList(left *ListNode, right *ListNode) *ListNode {
	tempList := &ListNode{}
	head := tempList
	for left != nil || right != nil {
		if right == nil || (left != nil && left.Val < right.Val) {
			tempList.Next = left
			left = left.Next
		} else {
			tempList.Next = right
			right = right.Next
		}
		tempList = tempList.Next
	}
	return head.Next
}

func main() {
	head := &ListNode{
		Val:  4,
		Next: nil,
	}
	head.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	head.Next.Next = &ListNode{
		Val:  1,
		Next: nil,
	}
	head.Next.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	sortList(head)
}
