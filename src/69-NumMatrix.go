package main

type NumMatrix struct {
	prefixMatrix [][]int
}

//func Constructor(matrix [][]int) NumMatrix {
//	rows, cols := len(matrix), len(matrix[0])
//	prefixMatrix := make([][]int, rows+1)
//	prefixMatrix[0] = make([]int, cols+1)
//	for i := 1; i < rows+1; i++ {
//		prefixMatrix[i] = make([]int, cols+1)
//		for j := 1; j < cols+1; j++ {
//			prefixMatrix[i][j] = matrix[i-1][j-1] + prefixMatrix[i-1][j] + prefixMatrix[i][j-1] - prefixMatrix[i-1][j-1]
//		}
//	}
//	return NumMatrix{prefixMatrix: prefixMatrix}
//}

func (this *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	return this.prefixMatrix[row2+1][col2+1] + this.prefixMatrix[row1][col1] - this.prefixMatrix[row1][col2+1] - this.prefixMatrix[row2+1][col1]
}
