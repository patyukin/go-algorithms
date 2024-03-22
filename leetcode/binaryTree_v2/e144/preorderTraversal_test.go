package e144

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	tests := []struct {
		name   string
		root   *TreeNode
		output []int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			output: []int{1, 2, 3},
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			output: []int{4, 2, 1, 3, 5, 6},
		},
		{
			root:   nil,
			output: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.root); !reflect.DeepEqual(got, tt.output) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversalLoop(tt.root); !reflect.DeepEqual(got, tt.output) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.output)
			}
		})
	}
}
