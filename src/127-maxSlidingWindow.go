package main

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	mq := NewMaxQueue()
	left, right := 0, 0
	for right < k {
		mq.push(nums[right])
		right++
	}
	res = append(res, mq.front())
	for right < len(nums) {
		mq.pop(nums[left])
		mq.push(nums[right])
		res = append(res, mq.front())
		left++
		right++
	}
	return res
}

type maxQueue struct {
	queue []int
}

func NewMaxQueue() *maxQueue {
	return &maxQueue{queue: []int{}}
}

func (mq *maxQueue) pop(num int) {
	if mq.front() == num {
		mq.queue = mq.queue[1:]
	}
}

func (mq *maxQueue) push(num int) {
	for len(mq.queue) > 0 && num > mq.back() {
		mq.queue = mq.queue[:len(mq.queue)-1]
	}
	mq.queue = append(mq.queue, num)
}

func (mq *maxQueue) front() int {
	return mq.queue[0]
}

func (mq *maxQueue) back() int {
	return mq.queue[len(mq.queue)-1]
}
