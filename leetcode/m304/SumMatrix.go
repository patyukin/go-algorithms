package m304

type NumMatrix struct {
	m [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rows := len(matrix)
	cols := len(matrix[0])

	m := make([][]int, rows+1)
	for i := range matrix {
		m[i] = make([]int, cols+1)
	}

	for row := 1; row <= rows; row++ {
		for col := 1; col <= cols; col++ {
			m[row][col] = m[row-1][col] + m[row][col-1] - m[row-1][col-1] + matrix[row-1][col-1]
		}
	}

	return NumMatrix{m: m}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row2 += 1
	col2 += 1

	return this.m[row2][col2] - this.m[row2][col1] - this.m[row1][col2] + this.m[row1][col1]
}
