package main

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	delMap := map[int]bool{}
	for _, delNode := range to_delete {
		delMap[delNode] = true
	}
	res := []*TreeNode{}
	var buildTree func(root *TreeNode) *TreeNode
	buildTree = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		root.Left = buildTree(root.Left)
		root.Right = buildTree(root.Right)
		if delMap[root.Val] {
			if root.Left != nil {
				res = append(res, root.Left)
			}
			if root.Right != nil {
				res = append(res, root.Right)
			}
			root = nil
		}
		return root
	}
	root = buildTree(root)
	if root != nil {
		res = append(res, root)
	}
	return res
}
