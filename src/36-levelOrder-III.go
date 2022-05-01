package main

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res, queue := [][]int{}, []*TreeNode{root}
	for len(queue) > 0 {
		temp := make([]int, len(queue))
		cnt := len(queue)
		// 偶数次 0、2... （其实0是第一次） 正序排列
		if len(res)&1 == 0 {
			for i := 0; i < cnt; i++ {
				temp[i] = queue[i].Val
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
		} else { // 逆序排列
			for i := 0; i < cnt; i++ {
				//temp[i] = queue[-i-1].Val	// slice不能用负数下标
				temp[i] = queue[cnt-1-i].Val
				// 以下的顺序还是按照正常顺序
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
		}
		res = append(res, temp)
		queue = queue[cnt:]
	}
	return res
}
func main() {

}
