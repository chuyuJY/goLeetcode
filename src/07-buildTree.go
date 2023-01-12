package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1. 递归算法
//func buildTree(preorder []int, inorder []int) *TreeNode { // 注意此处传进来的是slice
//	if len(preorder) == 0 { // 因为preorder 和 inorder 的长度是一样的，所以只判断一个即可
//		return nil
//	}
//	root := &TreeNode{Val: preorder[0]}
//	var index int
//	for ; index < len(inorder); index++ { // 找到左右子树的分界线，即根节点root在inorder中的index
//		if inorder[index] == preorder[0] {
//			break
//		}
//	}
//	root.Left = buildTree(preorder[1:index+1], inorder[:index]) // 要注意，preorder和inorder的数量是一样多的，不同的是根节点的位置，把根节点去除就行
//	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
//	return root // 当一个节点的左子树和右子树均赋值之后，递归给上一层
//}

var inorderMap map[int]int

// 构建哈希表优化遍历查找
// func buildTree(preorder []int, inorder []int) *TreeNode { // 注意此处传进来的是slice
// 	n := len(preorder)
// 	// 构造哈希表
// 	inorderMap = make(map[int]int, n)
// 	for i := 0; i < len(inorder); i++ {
// 		inorderMap[inorder[i]] = i
// 	}
// 	// 返回根节点
// 	return buildMyTree(preorder, inorder, 0, n-1, 0, n-1)

// }

// func buildMyTree(preorder, inorder []int, leftPreIndex, rightPreIndex, leftInIndex, rightInIndex int) *TreeNode {
// 	if leftInIndex > rightInIndex {
// 		return nil
// 	}
// 	root := &TreeNode{Val: preorder[leftPreIndex]}
// 	index := inorderMap[root.Val]
// 	// 得到左子树的长度
// 	length_left_subtree := index - leftInIndex
// 	root.Left = buildMyTree(preorder, inorder, leftPreIndex+1, leftPreIndex+length_left_subtree, leftInIndex, index-1)
// 	root.Right = buildMyTree(preorder, inorder, leftPreIndex+length_left_subtree+1, rightPreIndex, index+1, rightInIndex)
// 	return root
// }

// 2.
