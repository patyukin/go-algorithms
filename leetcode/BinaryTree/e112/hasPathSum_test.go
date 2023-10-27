package e112

import "testing"

func Test_hasPathSumSimple(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		sum  int
		want bool
	}{
		{
			name: "Test unsafeSlice",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			sum:  22,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.root, tt.sum); got != tt.want {
				t.Errorf("hasPathSumSimple() = %v, want %v", got, tt.want)
			}

			if got := hasPathSumSimple(tt.root, tt.sum); got != tt.want {
				t.Errorf("hasPathSumSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}
