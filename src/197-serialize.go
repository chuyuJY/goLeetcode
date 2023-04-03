package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// 先序遍历
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "N"
	}
	lists := []string{strconv.Itoa(root.Val)}
	lists = append(lists, this.serialize(root.Left))
	lists = append(lists, this.serialize(root.Right))
	return strings.Join(lists, ",")
}

//func (this *Codec) deserialize(data string) *TreeNode {
//	lists := strings.Split(data, ",")
//	i := 0
//	return this.bfs(lists, &i)
//}

//func (this *Codec) bfs(lists []string, i *int) *TreeNode {
//	cur := lists[*i]
//	*i++
//	if cur == "N" {
//		return nil
//	}
//	root := &TreeNode{}
//	root.Val, _ = strconv.Atoi(cur)
//	root.Left = this.bfs(lists, i)
//	root.Right = this.bfs(lists, i)
//	return root
//}

func (this *Codec) deserialize(data string) *TreeNode {
	lists := strings.Split(data, ",")
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		cur := lists[0]
		lists = lists[1:]
		if cur == "N" {
			return nil
		}
		root := &TreeNode{}
		root.Val, _ = strconv.Atoi(cur)
		root.Left = dfs()
		root.Right = dfs()
		return root
	}
	return dfs()
}
