package main

type LRUCache struct {
	capacity   int
	hashMap    map[int]*Node
	head, tail *Node
}

type Node struct {
	key, val  int
	pre, next *Node
}

//func Constructor(capacity int) LRUCache {
//	head, tail := &Node{}, &Node{}
//	head.next, tail.pre = tail, head
//	return LRUCache{
//		capacity: capacity,
//		hashMap:  make(map[int]*Node, capacity),
//		head:     head,
//		tail:     tail,
//	}
//}

func (this *LRUCache) Get(key int) int {
	// 1. 若存在此key
	////   1. 把该节点放到链表头部第一个，相当于更新
	////   2. 返回val
	if node, ok := this.hashMap[key]; ok {
		this.moveToHead(node)
		return node.val
	} else {
		// 2. 不存在该key
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	// 1. 若存在该key
	if node, ok := this.hashMap[key]; ok {
		// 改值
		node.val = value
		// 更新
		this.moveToHead(node)
		return
	}
	// 2. 若不存在
	// 若超过容量
	if this.capacity == len(this.hashMap) {
		rmKey := this.deleteTailNode()
		delete(this.hashMap, rmKey)
	}
	newNode := &Node{
		key:  key,
		val:  value,
		pre:  nil,
		next: nil,
	}

	this.addToHead(newNode)
	// 添加到hashMap
	this.hashMap[key] = newNode
}

// 移动到头部，更新
func (this *LRUCache) moveToHead(node *Node) {
	this.deleteNode(node)
	this.addToHead(node)
}

// 删除旧节点
func (this *LRUCache) deleteNode(node *Node) {
	node.pre.next, node.next.pre = node.next, node.pre
}

// 添加新的节点
func (this *LRUCache) addToHead(node *Node) {
	node.next, node.pre = this.head.next, this.head
	this.head.next, node.next.pre = node, node
}

// 删除多余节点
func (this *LRUCache) deleteTailNode() int {
	tailNode := this.tail.pre
	this.deleteNode(tailNode)
	return tailNode.key
}
