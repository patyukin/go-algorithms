package e101

import (
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	n7 := &TreeNode{Val: 0}
	n4 := &TreeNode{Val: 4}
	n3 := &TreeNode{Val: 3}
	n1 := &TreeNode{3, n3, n4}
	n2 := &TreeNode{3, n4, n3}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{1, n2, n5}

	tests := []struct {
		name string
		tree *TreeNode
		want bool
	}{
		{
			name: "Symmetric tree",
			tree: &TreeNode{1, n1, n2},
			want: true,
		},
		{
			name: "Non symmetric tree",
			tree: n6,
			want: false,
		},
		{
			name: "Tree with single node",
			tree: &TreeNode{Val: 1},
			want: true,
		},
		{
			name: "Empty tree",
			tree: &TreeNode{1, n7, nil},
			want: false,
		},
	}

	for _, c := range tests {
		if got := isSymmetric(c.tree); got != c.want {
			t.Errorf("test '%s' failed: got %v, want %v", c.name, got, c.want)
		}

		if got := isSymmetricRecursion(c.tree); got != c.want {
			t.Errorf("test '%s' failed: got %v, want %v", c.name, got, c.want)
		}

		if got := isSymmetricLoop(c.tree); got != c.want {
			t.Errorf("test '%s' failed: got %v, want %v", c.name, got, c.want)
		}
	}
}
