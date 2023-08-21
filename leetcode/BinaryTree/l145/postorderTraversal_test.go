package l145

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	tests := []struct {
		name   string
		root   *TreeNode
		expect []int
	}{
		{
			name:   "Test Case 1",
			root:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			expect: []int{2, 3, 1},
		},
		{
			name:   "Test Case 2",
			root:   nil,
			expect: []int{},
		},
		{
			name:   "Test Case 3",
			root:   &TreeNode{Val: 1},
			expect: []int{1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := postorderTraversal(tc.root); !reflect.DeepEqual(got, tc.expect) {
				t.Fatalf("postorderTraversal() = %v, want %v", got, tc.expect)
			}

			if got := postorderTraversalFast(tc.root); !reflect.DeepEqual(got, tc.expect) {
				t.Fatalf("postorderTraversal() = %v, want %v", got, tc.expect)
			}

			if got := postorderTraversalLoop(tc.root); !reflect.DeepEqual(got, tc.expect) {
				t.Fatalf("postorderTraversal() = %v, want %v", got, tc.expect)
			}
		})
	}
}
