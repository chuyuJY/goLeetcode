package main

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{queue: []int{}}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	cnt := len(this.queue) - 1
	for cnt > 0 {
		num := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, num)
		cnt--
	}
	num := this.queue[0]
	this.queue = this.queue[1:]
	return num
}

func (this *MyStack) Top() int {
	num := this.Pop()
	this.queue = append(this.queue, num)
	return num
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
