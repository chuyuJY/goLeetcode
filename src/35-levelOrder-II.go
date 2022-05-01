package main

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res, queue := [][]int{}, []*TreeNode{root}
	for len(queue) > 0 {
		temp := []int{}
		// 循环层数为queue当前节点数
		cnt := len(queue)
		for i := 0; i < cnt; i++ {
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

		}
		res = append(res, temp)
		queue = queue[cnt:]
	}
	return res
}
