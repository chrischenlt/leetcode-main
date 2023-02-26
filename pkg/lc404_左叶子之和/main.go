package lc404_左叶子之和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftValue := sumOfLeftLeaves(root.Left)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue = root.Left.Val
	}
	rightValue := sumOfLeftLeaves(root.Right)
	return leftValue + rightValue
}

var res int

func sumOfLeftLeaves1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res = 0
	a(root.Left, true)
	a(root.Right, false)
	return res
}
func a(root *TreeNode, tag bool) {
	if root == nil {
		return
	}
	if tag && root.Left == nil && root.Right == nil {
		res += root.Val
		return
	}

	a(root.Left, true)
	a(root.Right, false)
}
