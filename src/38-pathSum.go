package main

var res [][]int
var path []int

func pathSum(root *TreeNode, target int) [][]int {
	res = [][]int(nil)
	path = []int(nil)
	recur(root, target)
	return res
}
func recur(root *TreeNode, resTarget int) {
	if root == nil {
		return
	}
	// 递归公式就是如下：
	path = append(path, root.Val)
	resTarget -= root.Val
	// 递归到子节点，进行判断
	if resTarget == 0 && root.Left == nil && root.Right == nil {
		// 尤其注意此处！！！！！！！！！
		// 记录路径时若直接执行 res.append(path) ，则是将 path 对象加入了 res ；
		// 后续 path 改变时， res 中的 path 对象也会随之改变。
		res = append(res, append([]int(nil), path...))
		// 或这样写
		//temp := make([]int, len(path))
		//for i := 0; i < len(path); i++ {
		//	temp[i] = path[i]
		//}
		//res = append(res, temp)
	}
	recur(root.Left, resTarget)
	recur(root.Right, resTarget)
	path = path[:len(path)-1]
}
func main() {

}
