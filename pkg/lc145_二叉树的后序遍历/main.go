package lc145_二叉树的后序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 非递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}

	reverse(res)

	return res
}

func reverse(nums []int) {
	if len(nums) == 0 {
		return
	}

	start, end := 0, len(nums)-1
	for start <= end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
