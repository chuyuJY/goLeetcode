package main

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	graph := map[int][]int{}
	buildGraph(root, &graph)
	// bfs搜索
	visited := map[int]bool{}
	queue := []int{}
	nextQueue := []int{}
	copy(queue, graph[target.Val])
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		for _, nextNode := range graph[curNode] {
			if !visited[nextNode] {
				nextQueue = append(nextQueue, nextNode)
			}
		}
		if len(queue) == 0 {
			k--
			if k == 0 {
				return nextQueue
			}
			queue = nextQueue
			nextQueue = []int{}
		}
	}
	return []int{}
}

// func buildGraph(root *TreeNode, graph *map[int][]int) {
// 	if root == nil {
// 		return
// 	}
// 	if root.Left != nil {
// 		(*graph)[root.Val] = append((*graph)[root.Val], root.Left.Val)
// 		(*graph)[root.Left.Val] = append((*graph)[root.Left.Val], root.Val)
// 	}
// 	if root.Right != nil {
// 		(*graph)[root.Val] = append((*graph)[root.Val], root.Right.Val)
// 		(*graph)[root.Right.Val] = append((*graph)[root.Right.Val], root.Val)
// 	}
// 	buildGraph(root.Left, graph)
// 	buildGraph(root.Right, graph)
// }
