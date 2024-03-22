package e724

import (
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Пустой массив",
			nums: []int{},
			want: -1,
		},
		{
			name: "Единственный элемент в массиве",
			nums: []int{1},
			want: 0,
		},
		{
			name: "Нет индекса оси вращения",
			nums: []int{1, 2, 3},
			want: -1,
		},
		{
			name: "Несколько элементов, индекс оси вращения есть",
			nums: []int{1, 7, 3, 6, 5, 6},
			want: 3,
		},
		{
			name: "Несколько элементов, положительные и отрицательные числа",
			nums: []int{1, 7, 3, 6, 5, 6},
			want: 3,
		},
		{
			name: "Еще один случай",
			nums: []int{2, 1, -1},
			want: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := pivotIndex(tc.nums)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
