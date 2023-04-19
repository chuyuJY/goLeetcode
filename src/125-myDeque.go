package main

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type MyQueue struct {
	stack []int
	back  []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack: []int{},
		back:  []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.back) == 0 {
		for len(this.stack) > 0 {
			this.back = append(this.back, this.stack[len(this.stack)-1])
			this.stack = this.stack[:len(this.stack)-1]
		}
	}
	num := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return num
}

func (this *MyQueue) Peek() int {
	if len(this.back) == 0 {
		for len(this.stack) > 0 {
			this.back = append(this.back, this.stack[len(this.stack)-1])
			this.stack = this.stack[:len(this.stack)-1]
		}
	}
	return this.back[len(this.back)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.stack) == 0 && len(this.back) == 0 {
		return true
	}
	return false
}
