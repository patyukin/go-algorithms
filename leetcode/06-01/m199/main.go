package main

// 199. Binary Tree Right Side View
// https://leetcode.com/problems/binary-tree-right-side-view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if i == size-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]
	}

	return result
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var dfs func(node *TreeNode, level int, result *[]int)
	dfs = func(node *TreeNode, level int, result *[]int) {
		if node == nil {
			return
		}

		if level == len(*result) {
			*result = append(*result, node.Val)
		}

		dfs(node.Right, level+1, result)
		dfs(node.Left, level+1, result)
	}

	var result []int
	dfs(root, 0, &result)

	return result
}
