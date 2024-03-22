package m107

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrder(node *TreeNode, level int, levels *[][]int) {
	if len((*levels)[level]) == 0 {
		(*levels)[level] = []int{}
	}

	(*levels)[level] = append((*levels)[level], node.Val)
	levelOrder(node.Left, level+1, levels)
	levelOrder(node.Right, level+1, levels)
}

func reverseSlice(s *[][]int) {
	left, right := 0, len(*s)-1
	for left < right {
		(*s)[left], (*s)[right] = (*s)[right], (*s)[left]
		left++
		right--
	}
}

func levelOrderBottom(root *TreeNode) [][]int {
	var levels [][]int
	levelOrder(root, 0, &levels)
	reverseSlice(&levels)

	return levels
}
