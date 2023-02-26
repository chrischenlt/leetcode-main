package lc257_二叉树的所有路径

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []string

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res = make([]string, 0)

	recursion(root, "")

	return res
}

func recursion(root *TreeNode, str string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, fmt.Sprint(str, root.Val))
		return
	}
	str = fmt.Sprint(str, root.Val, "->")

	recursion(root.Left, str)
	recursion(root.Right, str)
}
