package m304

type NumMatrix struct {
	m [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	l := len(matrix)
	h := len(matrix[0])

	mx := make([][]int, l+1)
	for i := 0; i <= l; i++ {
		mx[i] = make([]int, h+1)
	}

	for i := 1; i <= l; i++ {
		for j := 1; j <= h; j++ {
			otherSum := mx[i-1][j] + mx[i][j-1] - mx[i-1][j-1]
			mx[i][j] = matrix[i-1][j-1] + otherSum
		}
	}

	return NumMatrix{m: mx}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	col2 += 1
	row2 += 1
	return this.m[row2][col2] - this.m[row2][col1] - this.m[row1][col2] + this.m[row1][col1]
}
