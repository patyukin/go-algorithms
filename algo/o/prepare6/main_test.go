package main

import (
	"reflect"
	"testing"
)

func Test_getPow(t *testing.T) {
	type args struct {
		n []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test case 1", args{[]int{1, 2, 3, 4, 5}}, []int{1, 4, 9, 16, 25}},
		{"test case 2", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}},
		{"test case 3", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}}, []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100, 121, 144, 169, 196, 225, 256, 289, 324, 361, 400}},
		{"test case 4", args{[]int{-3, -2, -1, 0, 1, 2, 3}}, []int{0, 1, 1, 4, 4, 9, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPow(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPow() = %v, want %v", got, tt.want)
			}

			if got := getPow2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPow2() = %v, want %v", got, tt.want)
			}
		})
	}
}
