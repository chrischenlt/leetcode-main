package lc226_翻转二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	recursion(root)
	return root
}

func recursion(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	recursion(root.Left)
	recursion(root.Right)
}
