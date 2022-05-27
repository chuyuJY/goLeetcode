package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, start, end int) *ListNode {
	if end == start {
		return lists[start]
	}
	mid := start + (end-start)/2
	return mergeList(merge(lists, start, mid), merge(lists, mid+1, end))
}

func mergeList(left, right *ListNode) *ListNode {
	tempList := &ListNode{}
	cur := tempList
	for left != nil || right != nil {
		if right == nil || (left != nil && left.Val < right.Val) {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	return tempList.Next
}
