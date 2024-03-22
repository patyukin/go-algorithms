package main

import (
	"reflect"
	"testing"
)

func Test_calc(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test case 1", args{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}}}, []int{2, 1, 4, 3, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calc(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
