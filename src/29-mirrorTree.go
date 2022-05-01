package main

//type TreeNode struct {
//	Val   int
//	left  *TreeNode
//	right *TreeNode
//}

func mirrorTree(root *TreeNode) *TreeNode {
	// 终止条件
	// 叶节点
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	root.Left = mirrorTree(root.Left)
	root.Right = mirrorTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
