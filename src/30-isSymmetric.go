package main

//func isSymmetric(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	return recur(root.Left, root.Right)
//}
//func recur(L, R *TreeNode) bool {
//	if L == nil && R == nil {
//		return true
//	} else if L == nil || R == nil || L.Val != R.Val {
//		return false
//	}
//	// 稍微复杂了点下面这种写法
//	//else if (L == nil && R != nil) || (L != nil && R == nil) {
//	//	return false
//	//}
//	return recur(L.Left, R.Right) && recur(L.Right, R.Left)
//}
