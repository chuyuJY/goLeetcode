package main

import "container/list"

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := list.New()
	if root == nil {
		return res
	}
	stack.PushBack(root)
	for stack.Len() > 0 {
		curNode := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, curNode.Val)
		if curNode.Left != nil {
			stack.PushBack(curNode.Left)
		}
		if curNode.Right != nil {
			stack.PushBack(curNode.Right)
		}
	}
	reverseList(res)
	return res
}

func reverseList(res []int) {
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
}
