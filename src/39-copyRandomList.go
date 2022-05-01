package main

//type Node struct {
//	Val    int
//	Next   *Node
//	Random *Node
//}

// 哈希表
//func copyRandomList(head *Node) *Node {
//	if head == nil {
//		return head
//	}
//	dic := map[*Node]*Node{}
//	cur := head
//	for ; cur != nil; cur = cur.Next {
//		temp := Node{cur.Val, cur.Next, cur.Random}
//		dic[cur] = &temp
//	}
//	cur = head
//	for ; dic[cur] != nil; cur = cur.Next {
//		dic[cur].Next = dic[cur.Next]
//		dic[cur].Random = dic[cur.Random]
//	}
//	return dic[head]
//}
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	// 建立新旧合并的链表
	pre, cur := head, head.Next
	for pre != nil {
		temp := Node{pre.Val, pre.Next, pre.Random}
		pre.Next = &temp
		// pre为最后一个节点时，cur为nil
		if cur == nil {
			break
		}
		pre, cur = cur, cur.Next
	}
	// 定义Random节点
	pre, cur = head, head.Next
	for pre != nil {
		if pre.Random != nil {
			cur.Random = pre.Random.Next
		} else {
			cur.Random = nil
		}
		if cur.Next == nil {
			break
		}
		pre, cur = pre.Next.Next, cur.Next.Next
	}
	// 分割开
	pre, cur = head, head.Next
	head = cur
	for pre != nil {
		pre.Next = pre.Next.Next
		if cur.Next == nil {
			break
		}
		cur.Next = cur.Next.Next
		pre, cur = pre.Next, cur.Next
	}
	return head
}
