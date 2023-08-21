package l94

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "TestCase1",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
		{
			name:     "TestCase2",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			expected: []int{2, 1},
		},
		{
			name:     "TestCase3",
			root:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			expected: []int{1, 2},
		},
		{
			name:     "TestCase4",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			expected: []int{2, 1, 3},
		},
		{
			name:     "TestCase5",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4}},
			expected: []int{3, 2, 1, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := inorderTraversal(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func Test_inorderTraversalLoop(t *testing.T) {
	type Test struct {
		name     string
		root     *TreeNode
		expected []int
	}

	tests := []Test{
		{
			name:     "test with empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "test with singleton tree",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
		{
			name: "test with complex tree",
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := inorderTraversalLoop(test.root)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func Test_inorderTraversalFaster(t *testing.T) {
	type Test struct {
		name     string
		root     *TreeNode
		expected []int
	}

	tests := []Test{
		{
			name:     "test with empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "test with singleton tree",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
		{
			name: "test with complex tree",
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := inorderTraversalFaster(test.root)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
