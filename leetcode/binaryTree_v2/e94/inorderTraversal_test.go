package e94

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Test with empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Test with singleton tree",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
		{
			name: "Test with complex tree",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	// Run test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := inorderTraversal(test.root)

			// Check if the result matches the expected output
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
