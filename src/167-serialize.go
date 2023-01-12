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

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		if curNode == nil {
			res = append(res, "N")
		} else {
			res = append(res, strconv.Itoa(curNode.Val))
			queue = append(queue, curNode.Left)
			queue = append(queue, curNode.Right)
		}
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "N" {
		return nil
	}
	lists := strings.Split(data, ",")
	val, _ := strconv.Atoi(lists[0])
	root := &TreeNode{
		Val: val,
	}
	queue := []*TreeNode{root}
	cursor := 1
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		if lists[cursor] != "N" {
			val, _ = strconv.Atoi(lists[cursor])
			curNode.Left = &TreeNode{
				Val: val,
			}
			queue = append(queue, curNode.Left)
		}
		if lists[cursor+1] != "N" {
			val, _ = strconv.Atoi(lists[cursor+1])
			curNode.Right = &TreeNode{
				Val: val,
			}
			queue = append(queue, curNode.Right)
		}
		cursor += 2
	}
	return root
}
