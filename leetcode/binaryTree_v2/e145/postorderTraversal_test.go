package e145

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Test Case 1",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			expected: []int{2, 3, 1},
		},
		{
			name:     "Test Case 2",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
			expected: []int{3, 2, 1},
		},
		{
			name:     "Test Case 3",
			root:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			expected: []int{2, 1},
		},
		{
			name:     "Test Case 4",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}},
			expected: []int{2, 4, 3, 1},
		},
		{
			name:     "Test Case 5",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
		{
			name: "Test Case 6",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
					Right: &TreeNode{Val: 5, Left: nil, Right: nil}},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
					Right: &TreeNode{Val: 7, Left: nil, Right: nil},
				},
			},
			expected: []int{4, 5, 2, 6, 7, 3, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := postorderTraversal(test.root)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := postorderTraversalLoop(test.root)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}

}
