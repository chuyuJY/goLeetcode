package main

import "github.com/emirpasic/gods/trees/redblacktree"

type TimeMap struct {
	hashMap map[string]*redblacktree.Tree
}

func Constructor() TimeMap {
	return TimeMap{
		hashMap: map[string]*redblacktree.Tree{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.hashMap[key]; !ok {
		this.hashMap[key] = redblacktree.NewWithIntComparator()
		this.hashMap[key].Put(-1, "") // 增加哨兵
	}
	this.hashMap[key].Put(timestamp, value)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if tree, ok := this.hashMap[key]; ok {
		floor, _ := tree.Floor(timestamp)
		return floor.Value.(string)
	}
	return ""
}
