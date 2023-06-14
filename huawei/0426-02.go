package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Val         int
	State       bool
	Left, Right *Node
}

type ResourcePool struct {
	head, tail *Node
	hashMap    map[int]*Node
}

func NewResourcePool(start, end int) *ResourcePool {
	rp := &ResourcePool{}
	for i := start; i <= end; i++ {
		node := &Node{
			Val: i,
		}
		rp.tail.Right, node.Left = node, rp.tail
		rp.tail = rp.tail.Right
		rp.hashMap[i] = node
	}
	return rp
}

func main() {
	sc := bufio.NewReader(os.Stdin)
	var start, end int
	fmt.Fscanln(sc, &start, &end)
	var op int
	fmt.Fscanln(sc, &op)
	opts := [][2]int{}
	for i := 0; i < op; i++ {
		var t, o int
		fmt.Fscanln(sc, &t, &o)
		opts = append(opts, [2]int{t, o})
	}
	getAnswer2(start, end, opts)
}

func getAnswer2(start, end int, opts [][2]int) {
	// 建立资源池
	rp := NewResourcePool(start, end)
	for _, opt := range opts {
		if opt[0] == 1 {
			// 如果还有空闲
			if rp.head.State {
				for opt[1] > 0 {

					rp.head = rp.head.Right
					rp.head.Left = nil
					opt[1]--
				}
			}
		} else if opt[0] == 2 {
			// 存在且没被占用
			if cur, exist := rp.hashMap[opt[1]]; exist {
				delete(rp.hashMap, cur.Val)
				cur.Left.Right, cur.Right.Left = cur.Right, cur.Left
			}
		} else if opt[0] == 3 {
			// 存在且被占用
			if _, exist := rp.hashMap[opt[1]]; !exist {
				node
			}
		}
	}
	fmt.Println(rp.head.Val)
}
