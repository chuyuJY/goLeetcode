package main

type SegNode struct {
	start int
	end   int
	left  *SegNode
	right *SegNode
}

type MyCalendar struct {
	root *SegNode
}

func Constructor() MyCalendar {
	return MyCalendar{root: &SegNode{start: -1, end: 0}} // 哨兵节点
}

func (this *MyCalendar) Book(start int, end int) bool {
	return Insert(this.root, &SegNode{start: start, end: end})
}

func Insert(root, node *SegNode) bool {
	// 插入左子树
	if node.end <= root.start {
		if root.left != nil {
			return Insert(root.left, node)
		}
		root.left = node
		return true
	}
	// 插入右子树
	if node.start >= root.end {
		if root.right != nil {
			return Insert(root.right, node)
		}
		root.right = node
		return true
	}
	return false
}
