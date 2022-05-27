package main

import (
	"fmt"
	"github.com/fatih/color"
	"sort"
)

// MaxMatch 记录最大匹配方案
type MaxMatch struct {
	min      int
	max      int
	n        int
	degrees  [][]int
	maxNodes []int
	result   []int
	resMax   int
}

// node 记录分配方案
type node struct {
	// 记录当前选择
	solution map[int]int
	// 记录当前的最大值
	maxSum int
}

// Construct MaxMatch 初始化参数
func Construct(degrees [][]int) MaxMatch {
	match := &MaxMatch{}
	match.degrees = degrees
	match.n = len(degrees)
	match.getLower()
	match.getUpper()
	match.result = make([]int, match.n)
	return *match
}

// MaxMatch 分支限界法求取最优分配
func (match *MaxMatch) bfsMatch() {
	// 队列实现bfs
	queue := []node{{
		// 物资-阵地编号
		solution: make(map[int]int),
		maxSum:   match.max,
	}}
	// 为每个阵地选择物资
	for i := 0; i < match.n; i++ {
		// 遍历每一层所有节点
		for cnt := len(queue); cnt > 0; cnt-- {
			// 给集装箱选择阵地
			for j := 0; j < match.n; j++ {
				// 队首元素
				temp := node{
					make(map[int]int),
					0,
				}
				for key, val := range queue[0].solution {
					temp.solution[key] = val
				}
				temp.maxSum = queue[0].maxSum
				// 判断该集装箱是否已分配
				if _, ok := temp.solution[j]; !ok {
					temp.solution[j] = i
					delt := match.degrees[i][j] - match.maxNodes[i]
					temp.maxSum += delt
					// 分支限界 剪枝
					if temp.maxSum >= match.min {
						queue = append(queue, temp)
					}
				}
			}
			// 出队
			queue = queue[1:]
		}
	}
	// 队列里剩余元素就是可分配的方案，排序选择最大的即可
	sort.Slice(queue, func(i, j int) bool {
		// 降序排列
		if queue[i].maxSum > queue[j].maxSum {
			return true
		} else {
			return false
		}
	})
	// key-物资, val-阵地
	for key, val := range queue[0].solution {
		// result的下标是阵地编号，值是物资编号
		match.result[val] = key
	}
	match.resMax = queue[0].maxSum
}

// MaxMatch 每行取最大，得到上界
func (match *MaxMatch) getUpper() {
	// 深复制
	temp := make([][]int, len(match.degrees))
	match.maxNodes = make([]int, match.n)
	for i := 0; i < len(match.degrees); i++ {
		for j := 0; j < len(match.degrees[i]); j++ {
			temp[i] = append(temp[i], match.degrees[i][j])
		}
	}
	// 获取每行最大值之和，并保存选项
	for i := 0; i < len(temp); i++ {
		sort.Ints(temp[i])
		tempMax := temp[i][len(temp[i])-1]
		match.max += tempMax
		match.maxNodes[i] = tempMax
	}
}

// MaxMatch 贪心法得到下界
func (match *MaxMatch) getLower() {
	// 用来标记对应列是否做过选择
	flag := make([]bool, match.n)
	// 贪心法遍历
	for i := 0; i < match.n; i++ {
		max := -1
		index := -1
		for j := 0; j < match.n; j++ {
			// 若未被选过，并且值更大
			if flag[j] == false && max < match.degrees[i][j] {
				max = match.degrees[i][j]
				index = j
			}
		}
		match.min += max
		flag[index] = true
	}
}

func main() {
	// 其中每个列表代表该阵地对物资的匹配值
	//matchDegree := [][]int{
	//	{10, 12, 20, 30, 18},
	//	{23, 30, 8, 12, 22},
	//	{11, 21, 23, 40, 16},
	//	{33, 34, 23, 19, 20},
	//	{21, 32, 11, 14, 21},
	//}
	//matchDegree := [][]int{
	//	{10, 12, 20},
	//	{23, 30, 8},
	//	{11, 21, 23},
	//}
	matchDegree := [][]int{
		{-1, -2, 0, -3},
		{-1, -2, -4, -3},
		{-2, -3, 0, -1},
		{-1, -1, 0, -2},
	}

	match := Construct(matchDegree)
	match.bfsMatch()
	fmt.Println(match.max)
	fmt.Println(match.min)
	fmt.Println(match.result)
	fmt.Println("  -------------------------------------------")
	color.Blue("  |             最佳物资投放方案            |")
	fmt.Println("  -------------------------------------------")
	for key, val := range match.result {
		fmt.Printf("  | 第%d号阵地投放编号%d 的物资--匹配度为: %d | \n", key+1, val+1, matchDegree[key][val])
	}
	fmt.Println("  -------------------------------------------")
	color.Green(fmt.Sprintf("  |           %s: %d           |", "最佳匹配度之和", match.resMax))
	fmt.Println("  -------------------------------------------")
}
