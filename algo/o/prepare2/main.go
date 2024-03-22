package main

// https://leetcode.com/problems/remove-nth-node-from-end-of-list

type ListNode struct {
	Val  int
	Next *ListNode
}

// Нужно удалить N-й элемент с конца.
func findAndDeleteFromEnd(root *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: root}
	first := dummy
	second := dummy

	for i := 1; i <= n+1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}
