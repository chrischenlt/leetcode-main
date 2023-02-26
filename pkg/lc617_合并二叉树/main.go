package lc617_合并二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	return &TreeNode{Val: root1.Val + root2.Val,
		Left: mergeTrees(root1.Left, root2.Left), Right: mergeTrees(root1.Right, root2.Right)}
}

//func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
//	if root1 == nil && root2 == nil {
//		return nil
//	}
//	var value int
//	var r1l, r1r, r2l, r2r *TreeNode
//	if root1 != nil {
//		value += root1.Val
//		r1l = root1.Left
//		r1r = root1.Right
//	}
//	if root2 != nil {
//		value += root2.Val
//		r2l = root2.Left
//		r2r = root2.Right
//	}
//
//	return &TreeNode{Val: value,
//		Left: mergeTrees(r1l, r2l), Right: mergeTrees(r1r, r2r)}
//}
