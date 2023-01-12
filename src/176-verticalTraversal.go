package main

import "sort"

type Position struct {
	Val    int
	Row    int
	Column int
}

func verticalTraversal(root *TreeNode) [][]int {
	positions := []Position{}
	inOrder(root, 0, 0, &positions)             // 传递数组指针
	sort.Slice(positions, func(i, j int) bool { // 排序
		if positions[i].Column == positions[j].Column {
			if positions[i].Row == positions[j].Row {
				return positions[i].Val < positions[j].Val
			}
			return positions[i].Row < positions[j].Row
		}
		return positions[i].Column < positions[j].Column
	})
	res := [][]int{{positions[0].Val}}
	for i := 1; i < len(positions); i++ {
		if positions[i].Column == positions[i-1].Column {
			res[len(res)-1] = append(res[len(res)-1], positions[i].Val)
		} else {
			res = append(res, []int{positions[i].Val})
		}
	}
	return res
}

func inOrder(root *TreeNode, row, column int, positions *[]Position) {
	if root == nil {
		return
	}
	*positions = append(*positions, Position{
		Val:    root.Val,
		Row:    row,
		Column: column,
	})
	inOrder(root.Left, row+1, column-1, positions)
	inOrder(root.Right, row+1, column+1, positions)
	return
}
