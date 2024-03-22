package m102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	customLevelOrder(root, &result, 0)

	return result
}

func customLevelOrder(root *TreeNode, preResult *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*preResult) == level {
		*preResult = append(*preResult, []int{})
	}

	(*preResult)[level] = append((*preResult)[level], root.Val)

	customLevelOrder(root.Left, preResult, level+1)
	customLevelOrder(root.Right, preResult, level+1)
}
