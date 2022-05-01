package main

/**
* Definition for singly-linked list.
 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//func reversePrint(head *ListNode) []int {
//	if head == nil {
//		return nil
//	}
//	return append(reversePrint(head.Next), head.Val)
//}

// 仔细体会，下面的方法是错误的！！！
//func reversePrint(head *ListNode) []int {
//	var l int
//	// 先遍历一遍 找出长度主要是
//	for head != nil {
//		head = head.Next
//		l++
//	}
//	reverse := make([]int, l)
//	for head != nil {
//		reverse[l-1] = head.Val // 注意不能用-l，因为数组不支持，slice支持负下标
//		head = head.Next
//		l--
//	}
//	return reverse
//}

// 原因是：你不能在第一个循环动head，这样指针的位置会变动，导致第二个循环出错！！！

func reversePrint(head *ListNode) []int {
	var l int
	//temp := head
	// 先遍历一遍 找出长度主要是
	//for head != nil {
	//	head = head.Next
	//	l++
	//}
	//head = temp

	// or

	for temp := head; temp != nil; temp = temp.Next {
		l++
	}
	reverse := make([]int, l)
	for head != nil {
		reverse[l-1] = head.Val // 注意不能用-l，因为数组不支持，slice支持负下标
		head = head.Next
		l--
	}
	return reverse
}

func main() {

}
