package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := map[*Node]*Node{}
	return dfsClone(node, visited)
}

func dfsClone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if v, ok := visited[node]; ok {
		return v
	}
	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = newNode
	for index, neighbor := range node.Neighbors {
		newNode.Neighbors[index] = dfsClone(neighbor, visited)
	}
	return newNode
}
