package main

import (
	"fmt"
	"time"
)

//
//import "fmt"
//
///*
//题目描述
//	给 n 个进程，每个进程都有一个独一无二的 PID （进程编号）和它的 PPID （父进程编号）。
//
//	每一个进程只有一个父进程，但是每个进程可能会有一个或者多个孩子进程，它们形成的关系就像一个树状结构。只有一个进程的 PPID 是 0 ，意味着这个进程没有父进程。所有的 PID 都会是唯一的正整数。
//
//	我们用两个序列来表示这些进程，第一个序列包含所有进程的 PID ，第二个序列包含所有进程对应的 PPID。
//
//	现在给定这两个序列和一个 PID 表示你要杀死的进程，函数返回一个 PID 序列，表示因为杀这个进程而导致的所有被杀掉的进程的编号。
//
//	当一个进程被杀掉的时候，它所有的孩子进程和后代进程都要被杀掉。
//	你可以以任意顺序排列返回的 PID 序列。
//
//示例1：
//输入:
//	pid = [1, 3, 10, 5]
//	ppid = [3, 0, 5, 3]
//	kill = 5
//输出: [5,10]
//解释:
//		3
//	   / \
//	  1   5
//		 /
//		10
//	杀掉进程 5 ，同时它的后代进程 10 也被杀掉。
//
//注意:
//	被杀掉的进程编号一定在 PID 序列中。
//	n >= 1.
//
//*/
//
//func main() {
//	pid := []int{1, 3, 10, 5}
//	ppid := []int{3, 0, 5, 3}
//	kill := 5
//	fmt.Println(test(pid, ppid, kill))
//}
//
//func test(pid []int, ppid []int, kill int) []int {
//	graph := map[int][]int{}
//	for i := 0; i < len(pid); i++ {
//		graph[ppid[i]] = append(graph[ppid[i]], pid[i])
//	}
//	res := []int{kill}
//	queue := []int{kill}
//	for len(queue) > 0 {
//		cur := queue[0]
//		queue = queue[1:]
//		for _, next := range graph[cur] {
//			queue = append(queue, next)
//			res = append(res, next)
//		}
//	}
//	return res
//}
//
//

//func main() {
//	ch := make(chan int)
//	//ch <- 1
//	close(ch)
//	for v := range ch {
//		fmt.Println(v)
//	}
//}

//func String2Bytes(s string) []byte {
//	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
//	bh := reflect.SliceHeader{
//		Data: sh.Data,
//		Len:  sh.Len,
//		Cap:  sh.Len,
//	}
//	return *(*[]byte)(unsafe.Pointer(&bh))
//}
//
//func Bytes2String(b []byte) string {
//	return *(*string)(unsafe.Pointer(&b))
//}

//type demo3 struct {
//	c int32
//	a struct{}
//}
//
//type demo4 struct {
//	a struct{}
//	c int32
//}
//
//func main() {
//	fmt.Println(unsafe.Sizeof(demo3{})) // 8
//	fmt.Println(unsafe.Sizeof(demo4{})) // 4
//}

var ch chan struct{}

func g1(ch1 chan int, ch2 chan int) {
	for i := 1; i <= 5; i++ {
		num := <-ch1
		fmt.Printf("%v%v", num, num+1)
		ch2 <- num
	}
}

func g2(ch1 chan int, ch2 chan int) {
	for i := 1; i <= 5; i++ {
		num := <-ch2
		fmt.Print(string('A'+num-1), string('A'+num))
		ch1 <- num + 2
	}
}

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	//ch = make(chan struct{}, 1)
	go g1(ch1, ch2)
	go g2(ch1, ch2)
	ch1 <- 1
	//<-ch
	time.Sleep(time.Second * 2)
}
