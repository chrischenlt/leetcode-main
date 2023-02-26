package lc110_平衡二叉树

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	height := getHeight(root)
	if height == -1 {
		return false
	}
	return true
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := getHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return -1
	} else {
		return 1 + max(leftHeight, rightHeight)
	}
}

//func isBalanced(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//
//	if !isBalanced(root.Left) {
//		return false
//	}
//
//	if !isBalanced(root.Right) {
//		return false
//	}
//
//	abs1 := math.Abs(float64(getHeight(root.Left) - getHeight(root.Right)))
//	if abs1 > 1 {
//		return false
//	}
//	return true
//}
//
//func getHeight(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	leftHeight := getHeight(root.Left)
//	rightHeight := getHeight(root.Right)
//	return max(leftHeight, rightHeight) + 1
//}
//
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
