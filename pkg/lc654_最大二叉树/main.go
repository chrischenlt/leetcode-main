package lc654_最大二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var index, maxV int
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxV {
			maxV = nums[i]
			index = i
		}
	}
	root := &TreeNode{Val: maxV}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}
