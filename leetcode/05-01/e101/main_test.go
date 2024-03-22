package main

import "testing"

func Test_isSymmetricTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Simple test",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}},
			want: true,
		},
		{
			name: "Simple test",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}},
			want: false,
		},
		{
			name: "Simple test",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}},
			want: false,
		},
		{
			name: "Simple test",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}},
			want: false,
		},
		{
			name: "Simple test",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetricTree(tt.args.root); got != tt.want {
				t.Errorf("isSymmetricTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
