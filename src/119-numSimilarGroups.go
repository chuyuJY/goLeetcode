package main

func numSimilarGroups(strs []string) int {
	fathers := make([]int, len(strs))
	for node := 0; node < len(strs); node++ {
		fathers[node] = node
	}

	res := len(strs)
	for nodeI := 0; nodeI < len(strs); nodeI++ {
		for nodeJ := nodeI + 1; nodeJ < len(strs); nodeJ++ {
			if isSimilar(strs[nodeI], strs[nodeJ]) && union(fathers, nodeI, nodeJ) {
				res--
			}
		}
	}
	return res
}

func isSimilar(strI, strJ string) bool {
	cnt := 0
	// 是否相似，其实就是二者不同的字符数 cnt 是否 <= 2 ?
	for i := 0; i < len(strI); i++ {
		if strI[i] != strJ[i] {
			cnt++
			if cnt > 2 {
				return false
			}
		}
	}
	return true
}

func findFather(fathers []int, node int) int {
	if fathers[node] != node {
		fathers[node] = findFather(fathers, fathers[node])
	}
	return fathers[node]
}

//func union(fathers []int, nodeI, nodeJ int) bool {
//	fatherI, fatherJ := findFather(fathers, nodeI), findFather(fathers, nodeJ)
//	if fatherI != fatherJ {
//		// 注意是更改根节点的father，易错点
//		fathers[fatherI] = fatherJ
//		return true
//	}
//	return false
//}
