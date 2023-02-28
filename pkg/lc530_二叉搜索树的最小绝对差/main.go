package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre *TreeNode
var res int

func getMinimumDifference(root *TreeNode) int {
	res = math.MaxInt64
	pre = nil
	a(root)
	return res
}
func a(root *TreeNode) {
	if root == nil {
		return
	}
	a(root.Left)
	if pre != nil {
		if res > root.Val-pre.Val {
			res = root.Val - pre.Val
		}
	}
	pre = root
	a(root.Right)
}

var arr []int

func getMinimumDifference2(root *TreeNode) int {
	arr = make([]int, 0)
	getArray(root)
	res := math.MaxInt64
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < res {
			res = arr[i] - arr[i-1]
		}
	}
	return res
}

func getArray(root *TreeNode) {
	if root == nil {
		return
	}
	getArray(root.Left)
	arr = append(arr, root.Val)
	getArray(root.Right)
}

//var res int
//
//func getMinimumDifference1(root *TreeNode) int {
//	res = math.MaxInt64
//	a(root, true)
//	a(root, false)
//	return res
//}
//
//func a(root *TreeNode, isLeft bool) {
//	if root == nil {
//		return
//	}
//	if isLeft {
//		cur := root.Left
//		if cur == nil {
//			return
//		}
//		for cur.Right != nil {
//			cur = cur.Right
//		}
//		if (root.Val - cur.Val) < res {
//			res = root.Val - cur.Val
//		}
//	} else {
//		cur := root.Right
//		if cur == nil {
//			return
//		}
//		for cur.Left != nil {
//			cur = cur.Left
//		}
//		if (cur.Val - root.Val) < res {
//			res = cur.Val - root.Val
//		}
//	}
//	a(root.Left, true)
//	a(root.Left, false)
//	a(root.Right, false)
//	a(root.Right, true)
//}
