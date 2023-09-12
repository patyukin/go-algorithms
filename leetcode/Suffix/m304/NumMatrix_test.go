package m304

import (
	"testing"
)

func TestNumMatrix_SumRegion(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		row1   int
		col1   int
		row2   int
		col2   int
		want   int
	}{
		{
			name:   "Test unsafeSlice: Small Matrix with positive numbers",
			matrix: [][]int{{1, 2}, {3, 4}},
			row1:   0, col1: 0, row2: 1, col2: 1,
			want: 10,
		},
		{
			name:   "Test reflectionSliceHeader: Small Matrix with negative numbers",
			matrix: [][]int{{-1, -2}, {-3, -4}},
			row1:   0, col1: 0, row2: 1, col2: 1,
			want: -10,
		},
		{
			name:   "Test 3: Single row",
			matrix: [][]int{{1, 2, 3, 4, 5}},
			row1:   0, col1: 0, row2: 0, col2: 4,
			want: 15,
		},
		{
			name:   "Test 4: Single column",
			matrix: [][]int{{1}, {2}, {3}, {4}, {5}},
			row1:   0, col1: 0, row2: 4, col2: 0,
			want: 15,
		},
		{
			name:   "Test 5: Large matrix",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			row1:   0, col1: 0, row2: 2, col2: 2,
			want: 45,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numMatrix := Constructor(tt.matrix)
			if got := numMatrix.SumRegion(tt.row1, tt.col1, tt.row2, tt.col2); got != tt.want {
				t.Errorf("NumMatrix.SumRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}
