package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := list.New()
	l.PushBack(root)
	res := 0
	for l.Len() > 0 {
		ll := l.Len()
		for i := 0; i < ll; i++ {
			front := l.Front()
			l.Remove(front)
			value := front.Value
			node := value.(*TreeNode)
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
		res++
	}
	return res
}
