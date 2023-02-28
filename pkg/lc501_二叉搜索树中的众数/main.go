package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//var arr []int
//
//func findMode(root *TreeNode) []int {
//	res := make([]int, 0)
//	arr = make([]int, 0)
//	A(root)
//	// 1 2 3 3 4
//	left, right := 0, 0
//	maxLength := 0
//	for right < len(arr) {
//		if arr[left] != arr[right] {
//			if right-left > maxLength {
//				res = []int{arr[left]}
//				maxLength = right - left
//
//			} else if right-left == maxLength {
//				res = append(res, arr[left])
//			}
//			left = right
//		}
//		right++
//	}
//	if right-left > maxLength {
//		res = []int{arr[left]}
//		left = right
//	} else if right-left == maxLength {
//		res = append(res, arr[left])
//	}
//	return res
//}
//
//func A(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	A(root.Left)
//	arr = append(arr, root.Val)
//	A(root.Right)
//}

var pre *TreeNode
var curLength, maxLength int
var res []int

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	pre = nil
	maxLength = 0
	curLength = 1
	res = make([]int, 0)

	a(root)
	return res
}

func a(root *TreeNode) {
	if root == nil {
		return
	}
	a(root.Left)
	if pre != nil {
		if pre.Val == root.Val {
			curLength++
		} else {
			curLength = 1
		}
	} else {
		curLength = 1
	}
	pre = root
	if curLength > maxLength {
		res = []int{pre.Val}
		maxLength = curLength
	} else if curLength == maxLength {
		res = append(res, pre.Val)
	}
	a(root.Right)
}
