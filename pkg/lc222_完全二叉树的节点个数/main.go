package lc222_完全二叉树的节点个数

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := root.Left, root.Right
	leftDepth, rightDepth := 0, 0
	for left != nil {
		left = left.Left
		leftDepth++
	}
	for right != nil {
		right = right.Right
		rightDepth++
	}
	if leftDepth == rightDepth {
		return 2<<leftDepth - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

var res int

func countNodes1(root *TreeNode) int {
	res = 0
	if root == nil {
		return res
	}

	list := []*TreeNode{root}
	for len(list) > 0 {
		a := list[0]
		list = list[1:]
		res++
		if a.Left != nil {
			list = append(list, a.Left)
		}
		if a.Right != nil {
			list = append(list, a.Right)
		}
	}

	return res
}

func countNodes2(root *TreeNode) int {
	res = 0
	recursion(root)

	return res
}

func recursion(root *TreeNode) {
	if root == nil {
		return
	}
	res++
	recursion(root.Left)
	recursion(root.Right)
}
