package l144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversalLoop(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, current.Val)
		if current.Right != nil {
			stack = append(stack, current.Right)
		}

		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return result
}

func preorderTraversalTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// Список для хранения значений узлов в порядке обхода
	var result []int

	// Вспомогательная функция для рекурсивного обхода дерева в порядке preorder
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Добавляем значение текущего узла в результат
		result = append(result, node.Val)
		// Рекурсивно обходим левое поддерево
		preorder(node.Left)
		// Рекурсивно обходим правое поддерево
		preorder(node.Right)
	}

	// Вызываем вспомогательную функцию preorder для выполнения обхода
	preorder(root)

	return result
}
