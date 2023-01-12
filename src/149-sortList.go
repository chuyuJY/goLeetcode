package main

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right := splitList(head)
	return mergeSort(left, right)
}

func splitList(head *ListNode) (*ListNode, *ListNode) {
	// 二分
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	left, right := head, slow.Next
	slow.Next = nil
	return left, right
}

func mergeSort(left, right *ListNode) *ListNode {
	// 归并
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
