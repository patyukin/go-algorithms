package m1314

func matrixBlockSum(mat [][]int, k int) [][]int {
	rows := len(mat)
	cols := len(mat[0])

	m := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		m[i] = make([]int, cols+1)
	}

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			m[i][j] = m[i-1][j] + m[i][j-1] - m[i-1][j-1] + mat[i-1][j-1]
		}
	}

	output := make([][]int, rows)
	for i := range output {
		output[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			r1 := max(0, i-k)
			r2 := min(rows, i+k+1)
			c1 := max(0, j-k)
			c2 := min(cols, j+k+1)

			output[i][j] = m[r2][c2] - m[r2][c1] - m[r1][c2] + m[r1][c1]
		}
	}

	return output
}

// Вспомогательная функция для определения максимума
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// Вспомогательная функция для определения минимума
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
