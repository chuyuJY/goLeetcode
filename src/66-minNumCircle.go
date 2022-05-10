package main

//type Node struct {
//	x float64
//	y float64
//	r float64
//}
//
//func minNumCircle(x, y []float64, d float64) ([]float64, int) {
//	length := len(x)
//	nodes := make([]Node, length)
//	for i := 0; i < length; i++ {
//		nodes[i].x = x[i]
//		nodes[i].y = y[i]
//		nodes[i].r = x[i] + math.Sqrt(math.Pow(d, 2)-math.Pow(y[i], 2))
//	}
//	// 排序
//	var quickCirSort func(nodes []Node, l, r int) []Node
//	quickCirSort = func(nodes []Node, l, r int) []Node {
//		if l >= r {
//			return nodes
//		}
//		i, j := l, r
//		for i < j {
//			for i < j && nodes[j].r >= nodes[l].r {
//				j--
//			}
//			for i < j && nodes[i].r <= nodes[l].r {
//				i++
//			}
//			nodes[i], nodes[j] = nodes[j], nodes[i]
//		}
//		nodes[i], nodes[l] = nodes[l], nodes[i]
//		quickCirSort(nodes, l, i-1)
//		quickCirSort(nodes, i+1, r)
//		return nodes
//	}
//	quickCirSort(nodes, 0, length-1)
//	// 记录圆心
//	cnt, flag := []float64{}, 0
//	for flag < length {
//		cnt = append(cnt, nodes[flag].r)
//		flag++
//		for flag < length && math.Sqrt(math.Pow(nodes[flag].x-cnt[len(cnt)-1], 2)+math.Pow(nodes[flag].y, 2)) <= d {
//			flag++
//		}
//	}
//	return cnt, len(cnt)
//}
//
//func main() {
//	x := []float64{1, 2, 2, 3, 5, 6, 7, 8, 9, 10, 11, 13, 15}
//	y := []float64{1, 3, 4, 5, 4, 3, 3, 2, 4, 2, 3, 4, 5}
//	circle, num := minNumCircle(x, y, 5)
//	fmt.Println(circle, num)
//}
