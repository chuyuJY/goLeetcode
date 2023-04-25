package main

import (
	"bufio"
	"fmt"
	"os"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	pre := []int{}
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(sc, &num)
		pre = append(pre, num)
	}
	fmt.Fscanln(sc)
	in := []int{}
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(sc, &num)
		in = append(in, num)
	}
	fmt.Fscanln(sc)
	test3(pre, in)
}

func test3(pre, in []int) {
	// 建树
	root := buildTree(pre, in)
	// 判断是否对称
	maxNode, res, maxAll := 0, 0, 0
	var isSymmetric func(root *TreeNode) bool
	isSymmetric = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if root.Val > maxAll {
			maxAll = root.Val
		}
		var compare func(left, right *TreeNode) bool
		compare = func(left, right *TreeNode) bool {
			if left == nil && right == nil {
				return true
			}
			if left != nil && right != nil {
				if left.Val > right.Val {
					if left.Val > maxNode {
						maxNode, res = left.Val, right.Val
					}
				} else {
					if right.Val > maxNode {
						maxNode, res = right.Val, left.Val
					}
				}
				return compare(left.Left, right.Right) && compare(left.Right, right.Left)
			}
			return false
		}
		return compare(root.Left, root.Right)
	}

	isValid := isSymmetric(root)
	if isValid {
		fmt.Println(res)
	} else {
		fmt.Println(maxAll)
	}
}

func buildTree(pre, in []int) *TreeNode {
	if len(pre) == 0 || len(in) == 0 {
		return nil
	}
	root := &TreeNode{Val: pre[0]}
	index := 0
	for index < len(in) {
		if in[index] == pre[0] {
			break
		}
		index++
	}
	if index == len(in) {
		index = -1
	}
	root.Left = buildTree(pre[1:index+1], in[:index])
	root.Right = buildTree(pre[index+1:], in[index+1:])
	return root
}
