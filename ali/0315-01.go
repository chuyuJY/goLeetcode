package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	nodes := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		var l, r int
		fmt.Fscanln(sc, &l, &r)
		nodes[i] = []int{l, r}
	}
	test3(n, nodes)
}

func test3(n int, nodes [][]int) {
	deeps := make([]int, n+1)
	var dfs func(node int)
	dfs = func(node int) {
		left, right := nodes[node][0], nodes[node][1]
		// 当前节点为叶节点
		if left == -1 && right == -1 {
			deeps[node] = 1
			return
		}
		// 当前节点只有左子树或右子树
		if left == -1 {
			deeps[node] = -1
			dfs(right)
			return
		}
		if right == -1 {
			deeps[node] = -1
			dfs(left)
			return
		}
		// 当前节点的左右子树还未搜索
		if deeps[left] == 0 {
			dfs(left)
		}
		if deeps[right] == 0 {
			dfs(right)
		}
		// 当前节点左右子树都搜索过了, 且左右子树不都为满二叉树
		if deeps[left] == -1 || deeps[right] == -1 {
			deeps[node] = -1
			return
		}
		// 当前节点的左右子树均为满二叉树, 且高度一致
		if deeps[left] == deeps[right] {
			deeps[node] = deeps[left] + 1
		} else {
			deeps[node] = -1
		}
		return
	}
	dfs(1)
	res := 0
	for i := 1; i < len(deeps); i++ {
		if deeps[i] > 0 {
			res++
		}
	}
	fmt.Println(res)
}
