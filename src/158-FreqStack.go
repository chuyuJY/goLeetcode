package main

type FreqStack struct {
	maxFre   int           // 当前最大频率
	freMap   map[int]int   // 统计频率
	stackMap map[int][]int // 统计同频率元素的出现顺序
}

func Constructor() FreqStack {
	return FreqStack{
		maxFre:   0,
		freMap:   map[int]int{},
		stackMap: map[int][]int{},
	}
}

func (this *FreqStack) Push(val int) {
	this.freMap[val]++
	fre := this.freMap[val]
	if fre > this.maxFre {
		this.maxFre = fre
	}
	this.stackMap[fre] = append(this.stackMap[fre], val)
}

func (this *FreqStack) Pop() int {
	length := len(this.stackMap[this.maxFre])
	val := this.stackMap[this.maxFre][length-1]
	// 弹出
	this.freMap[val]--
	this.stackMap[this.maxFre] = this.stackMap[this.maxFre][:length-1]
	if length-1 == 0 {
		this.maxFre--
	}
	return val
}
