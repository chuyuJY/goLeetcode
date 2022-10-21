package main

import "container/list"

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := list.New()
	if root == nil {
		return res
	}
	curNode := root
	for curNode != nil || stack.Len() > 0 {
		if curNode != nil {
			stack.PushBack(curNode)
			curNode = curNode.Left
		} else {
			curNode = stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, curNode.Val)
			curNode = curNode.Right
		}
	}
	return res
}
