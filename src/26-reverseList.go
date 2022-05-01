package main

//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	left := head
//	head = head.Next
//	left.Next = nil
//	for head != nil {
//		right := head.Next
//		head.Next = left
//		left = head
//		head = right
//	}
//	return left
//}
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return temp
}
