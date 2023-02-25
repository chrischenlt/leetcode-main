package lc111_二叉树的最小深度

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	list := []*TreeNode{root}
	depth := 1
	re := 10000
	for len(list) > 0 {
		length := len(list)
		newList := make([]*TreeNode, 0)
		for i := 0; i < length; i++ {

			a := list[0]
			list = list[1:]
			if a == nil {
				continue
			}
			if a.Left == nil && a.Right == nil {
				re = min(re, depth)
				continue
			}

			newList = append(newList, a.Left)
			newList = append(newList, a.Right)
		}
		list = newList

		depth++
	}

	return re
}

var res int

func minDepth1(root *TreeNode) int {
	res = 10000
	if root == nil {
		return 0
	}
	getMin(root, 1)
	return res
}

func getMin(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		res = min(res, depth)
	}

	getMin(root.Left, depth+1)
	getMin(root.Right, depth+1)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
