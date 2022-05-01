package main

//func isSubStructure(A *TreeNode, B *TreeNode) bool {
//	if A == nil || B == nil {
//		return false
//	}
//	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
//}
//func recur(A *TreeNode, B *TreeNode) bool {
//	// 终止条件
//	if B == nil {
//		return true
//	}
//	// 此时 B不为nil
//	if A == nil || A.Val != B.Val {
//		return false
//	}
//	// 此时，A.Val==B.Val，可以继续判断
//	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
//}
