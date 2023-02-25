package main

import "fmt"

func main() {
	node := TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	traversal := preorderTraversal(&node)
	fmt.Println(traversal)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 非递归方法
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return res
}

// 递归方法
func preorderTraversal1(root *TreeNode) []int {

	res := make([]int, 0, 100)
	a(root, &res)
	return res
}

// 不加指针会返回空数组，外层的数组的len未发生变化
func a(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	*nums = append(*nums, root.Val)
	a(root.Left, nums)
	a(root.Right, nums)
}
