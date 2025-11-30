type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rowSumPrefix := 0
	var sumRegion [][]int
	sumRegion = make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		sumRegion[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 {
				sumRegion[i][j] = rowSumPrefix + matrix[i][j]
			} else {
				sumRegion[i][j] = sumRegion[i-1][j] + rowSumPrefix + matrix[i][j]
			}
			rowSumPrefix += matrix[i][j]
		}
		rowSumPrefix = 0
	}

	return NumMatrix{
		matrix: sumRegion,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	S := this.matrix
	result := S[row2][col2]

	if row1 > 0 {
		result -= S[row1-1][col2]
	}

	if col1 > 0 {
		result -= S[row2][col1-1]
	}

	if row1 > 0 && col1 > 0 {
		result += S[row1-1][col1-1]
	}

	return result
}
