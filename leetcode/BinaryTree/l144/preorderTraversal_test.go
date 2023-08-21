package l144

import (
	"reflect"
	"testing"
)

func Test_preorderTraversalLoop(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "TestCase1",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "TestCase2",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "TestCase3",
			root:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "TestCase4",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "TestCase5",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}},
			expected: []int{1, 2, 4, 5, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := preorderTraversalLoop(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func Test_preorderTraversalTree(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "TestCase1",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "TestCase2",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "TestCase3",
			root:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "TestCase4",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "TestCase5",
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}},
			expected: []int{1, 2, 4, 5, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := preorderTraversalTree(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
