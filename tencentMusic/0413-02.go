package main

import (
	"bufio"
	"fmt"
	"os"
)

type TreeNode struct {
	Id          int
	Nums        []int // 存储 2 和 5 的个数, 直接存乘积可能会溢出
	Left, Right *TreeNode
}

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	a := []int{0}
	for i := 1; i <= n; i++ {
		var ai int
		fmt.Fscan(sc, &ai)
		a = append(a, ai)
	}
	fmt.Fscanln(sc)
	edges := [][]int{}
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscanln(sc, &u, &v)
		edges = append(edges, []int{u, v})
	}
	test2(a, edges)
}

func test2(a []int, edges [][]int) {
	// 1. 建树
	tree := make([]*TreeNode, len(a))
	for _, edge := range edges {
		if tree[edge[0]] == nil {
			tree[edge[0]] = &TreeNode{}
		}
		if tree[edge[1]] == nil {
			tree[edge[1]] = &TreeNode{}
		}
		tree[edge[0]].Id, tree[edge[1]].Id = edge[0], edge[1]
		tree[edge[0]].Nums = calZero(a[edge[0]])
		tree[edge[1]].Nums = calZero(a[edge[1]])
		if tree[edge[0]].Left == nil {
			tree[edge[0]].Left = tree[edge[1]]
		} else {
			tree[edge[0]].Right = tree[edge[1]]
		}
	}
	root := tree[1]
	// 2. 后序遍历
	res := make([]int, len(a))
	var sufOrder func(curNode *TreeNode)
	sufOrder = func(curNode *TreeNode) {
		if curNode.Left == nil && curNode.Right == nil {
			res[curNode.Id] = min(curNode.Nums[0], curNode.Nums[1])
			return
		}
		if curNode.Left != nil {
			sufOrder(curNode.Left)
			curNode.Nums[0] += curNode.Left.Nums[0]
			curNode.Nums[1] += curNode.Left.Nums[1]
		}
		if curNode.Right != nil {
			sufOrder(curNode.Right)
			curNode.Nums[0] += curNode.Right.Nums[0]
			curNode.Nums[1] += curNode.Right.Nums[1]
		}
		res[curNode.Id] = min(curNode.Nums[0], curNode.Nums[1])
	}
	sufOrder(root)
	for i := 1; i < len(res); i++ {
		fmt.Printf("%v ", res[i])
	}
}

func calZero(val int) []int {
	nums := make([]int, 2)
	for val > 0 && val%2 == 0 {
		nums[0]++
		val /= 2
	}
	for val > 0 && val%5 == 0 {
		nums[1]++
		val /= 5
	}
	return nums
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
