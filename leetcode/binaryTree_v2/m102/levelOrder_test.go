package m102

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: [][]int{{1}, {2, 3}, {4, 5}},
		},
		{
			name:     "Example 3",
			root:     (*TreeNode)(nil),
			expected: [][]int{},
		},
		{
			name: "Tree with only one node",
			root: &TreeNode{
				Val: 1,
			},
			expected: [][]int{{1}},
		},
		{
			name: "Tree with multiple levels and only left children",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
					},
				},
			},
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name: "Tree with multiple levels and only right children",
			root: &TreeNode{
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
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name: "Tree with a single branch of left or right children",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
					},
				},
			},
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := levelOrder(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
