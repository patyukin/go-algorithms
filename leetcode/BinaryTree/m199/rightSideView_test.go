package m199

import (
	"testing"
)

func Test_rightSideView(t *testing.T) {
	tests := []struct {
		name     string
		tree     *TreeNode
		expected []int
	}{
		{
			name: "Test Case 1",
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			expected: []int{1, 3, 4},
		},
		{
			name: "Test Case 2",
			tree: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
			},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Test Case 3",
			tree:     nil,
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := rightSideView(test.tree)
			if !isEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}

func isEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
