//package main
//
//type MapSum struct {
//	next [26]*MapSum
//	val  int
//}
//
//func Constructor() MapSum {
//	return MapSum{
//		next: [26]*MapSum{},
//		val:  0,
//	}
//}
//
//func (this *MapSum) Insert(key string, val int) {
//	cur := this
//	for _, char := range key {
//		index := char - 'a'
//		if cur.next[index] == nil {
//			cur.next[index] = &MapSum{
//				next: [26]*MapSum{},
//				val:  0,
//			}
//		}
//		cur = cur.next[index]
//	}
//	// 更新该节点的val
//	cur.val = val
//}
//
//func (this *MapSum) Sum(prefix string) int {
//	cur := this
//	for _, char := range prefix {
//		index := char - 'a'
//		if cur.next[index] == nil {
//			return 0
//		}
//		cur = cur.next[index]
//	}
//	return dfsGetSum(cur)
//}
//
//func dfsGetSum(root *MapSum) int {
//	// 终止条件
//	if root == nil {
//		return 0
//	}
//	res := root.val
//	for i := 0; i < 26; i++ {
//		res += dfsGetSum(root.next[i])
//	}
//	return res
//}
