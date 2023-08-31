package m1314

import (
	"reflect"
	"testing"
)

func Test_matrixBlockSum(t *testing.T) {
	var MatrixBlockSumTests = []struct {
		name     string
		matrix   [][]int
		k        int
		expected [][]int
	}{
		{
			name:     "Тест 1: Базовый сценарий",
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:        1,
			expected: [][]int{{12, 21, 16}, {27, 45, 33}, {24, 39, 28}},
		},
		{
			name:     "Тест 2: Матрица 1x1",
			matrix:   [][]int{{1}},
			k:        1,
			expected: [][]int{{1}},
		},
		{
			name:     "Тест 3: Большое окно",
			matrix:   [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			k:        2,
			expected: [][]int{{4, 6, 4}, {6, 9, 6}, {4, 6, 4}},
		},
		{
			name:     "Тест 4: Нулевое окно",
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:        0,
			expected: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			name:     "Тест 4: Нулевое окно",
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:        1,
			expected: [][]int{{12, 21, 16}, {27, 45, 33}, {24, 39, 28}},
		},
		{
			name:     "Тест 5: Отрицательные числа",
			matrix:   [][]int{{-1, -2, -3}, {-4, -5, -6}, {-7, -8, -9}},
			k:        1,
			expected: [][]int{{-12, -21, -16}, {-27, -45, -33}, {-24, -39, -28}},
		},
	}

	for _, tt := range MatrixBlockSumTests {
		t.Run(tt.name, func(t *testing.T) {
			result := matrixBlockSum(tt.matrix, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v but got %v", tt.expected, result)
			}
		})
	}
}
