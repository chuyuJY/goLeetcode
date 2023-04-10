package main

import "fmt"

// 将二叉搜索树转化为双向链表

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeToList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var head, tail *TreeNode
	var inOrder func(cur *TreeNode)
	inOrder = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		inOrder(cur.Left)
		if tail != nil {
			tail.Right = cur
			cur.Left = tail
		} else {
			head = cur
		}
		tail = cur
		inOrder(cur.Right)
	}
	inOrder(root)
	head.Left, tail.Right = tail, head
	return head
}

func main() {
	root := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{5, nil, nil}}
	res := treeToList(root)
	for i := 0; i < 5; i++ {
		fmt.Println(res.Val)
		res = res.Right
	}
}
