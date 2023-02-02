package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	a(root)
	return res
}

var res int

func a(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left, right := a(node.Left), a(node.Right)
	if left+right > res {
		res = left + right
	}

	if left >= right {
		return left + 1
	}
	return right + 1
}
