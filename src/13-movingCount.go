package main

import "fmt"

// 定义两个标记，一个用来记录可达节点数，一个用来标记已到达过的节点
var num int
var visted [][]bool

func movingCount(m int, n int, k int) int {
	// 对二维数组进行分配空间
	visted = make([][]bool, m)
	for i := 0; i < m; i++ {
		visted[i] = make([]bool, n)
	}
	return dfs1(m, n, 0, 0, k)
}

// dsf函数
func dfs1(m, n, i, j, k int) int {
	// 递归结束条件
	//失败的条件：越界、到达过、数位和不符合条件
	if i >= m || j >= n || visted[i][j] || sum(i)+sum(j) > k {
		return 0
	}
	// 成功，进行标记
	visted[i][j] = true
	// 通过这个点，再向下和右查找，其实足够遍历到所有的可达节点了，无需向上下左右都进行查找（上面那道题是因为想要的是路径匹配，这个题只需要找出能经过的即可）
	num = 1 + dfs1(m, n, i+1, j, k) + dfs1(m, n, i, j+1, k)
	return num
}

// 求数位和
func sum(x int) (s int) {
	for ; x != 0; x /= 10 {
		s += x % 10
	}
	return
}

func main() {
	fmt.Println(sum(666))
}
