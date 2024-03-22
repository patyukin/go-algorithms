package main

import "testing"

func Test_maxNumberProductWithSort(t *testing.T) {
	type args struct {
		s []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{[]int{1, 2, 3, 4, 5}}, 20},
		{"test 2", args{[]int{1, 2, 3, 4, 5, 6}}, 30},
		{"test 3", args{[]int{1, 2, 3, 4, 5, 6, 7}}, 42},
		{"test 4", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}}, 56},
		{"test 4", args{[]int{1, 2, 3, 4, 5, 6, -7, -8}}, 56},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberProductWithSort(tt.args.s); got != tt.want {
				t.Errorf("maxNumberProductWithSort() = %v, want %v", got, tt.want)
			}

			if got := maxNumberProductWithoutSort(tt.args.s); got != tt.want {
				t.Errorf("maxNumberProductWithoutSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
