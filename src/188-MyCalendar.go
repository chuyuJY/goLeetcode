package main

import "github.com/emirpasic/gods/trees/redblacktree"

type MyCalendar struct {
	Tree *redblacktree.Tree
}

func Constructor() MyCalendar {
	tree := redblacktree.NewWithIntComparator()
	tree.Put(-1, -1) // 增加哨兵
	return MyCalendar{Tree: tree}
}

func (this *MyCalendar) Book(start int, end int) bool {
	floor, _ := this.Tree.Floor(start)
	if floor.Value.(int) > start {
		return false
	}
	if it := this.Tree.IteratorAt(floor); it.Next() && it.Key().(int) < end {
		return false
	}
	this.Tree.Put(start, end)
	return true
}
