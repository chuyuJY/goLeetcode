package main

import "container/list"

//func levelOrder(root *TreeNode) []int {
//	var temp []int
//	var queue []*TreeNode
//	if root == nil {
//		return temp
//	}
//	queue = append(queue, root)
//	for len(queue) > 0 {
//		temp = append(temp, queue[0].Val)
//		if queue[0].Left != nil {
//			queue = append(queue, queue[0].Left)
//		}
//		if queue[0].Right != nil {
//			queue = append(queue, queue[0].Right)
//		}
//		queue = queue[1:]
//	}
//	return temp
//}

// 用list做一遍！
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var temp []int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		// 强制转换类型，其实相当于声明这个类型。不然无法访问内部结构
		subroot := queue.Front().Value.(*TreeNode)
		temp = append(temp, subroot.Val)
		if subroot.Left != nil {
			queue.PushBack(subroot.Left)
		}
		if subroot.Right != nil {
			queue.PushBack(subroot.Right)
		}
		queue.Remove(queue.Front())
	}
	return temp
}
func main() {

}
