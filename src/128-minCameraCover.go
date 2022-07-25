package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	res := 0
	// 定义三种状态：0，1，2，分别代表无覆盖，有摄像头，有覆盖
	// 后序遍历
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		// 叶节点的空子节点
		if root == nil {
			return 2
		}

		left, right := dfs(root.Left), dfs(root.Right)
		// 注意顺序不能换
		if left == 2 && right == 2 {
			return 0
		}
		if left == 0 || right == 0 {
			res++
			return 1
		}
		if left == 1 || right == 1 {
			return 2
		}

		// 不会执行到此
		return -1
	}
	if dfs(root) == 0 {
		res++
	}
	return res
}
