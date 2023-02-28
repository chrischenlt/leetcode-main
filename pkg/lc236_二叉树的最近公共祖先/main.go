package main

import "fmt"

func main() {
	a := []int{1, 2}
	b := a

	a[0] = 2
	fmt.Println(a)
	fmt.Println(b)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)
	if leftNode != nil && rightNode != nil {
		return root
	}
	if leftNode != nil && rightNode == nil {
		return leftNode
	}
	if leftNode == nil && rightNode != nil {
		return rightNode
	}
	return nil
}

var parr, qarr []int

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	way := make([]int, 0)
	parr = make([]int, 0)
	qarr = make([]int, 0)
	traversal(root, p.Val, q.Val, way)
	var long, short []int
	if len(parr) > len(qarr) {
		long = parr
		short = qarr
	} else {
		long = qarr
		short = parr
	}
	for i := len(long) - 1; i >= 0; i-- {
		tag := long[i]
		for i2 := len(short) - 1; i2 >= 0; i2-- {
			if tag == short[i2] {
				return &TreeNode{Val: tag}
			}
		}
	}
	return nil
}

func traversal(root *TreeNode, p, q int, way []int) {
	if root == nil || (len(parr) > 0 && len(qarr) > 0) {
		return
	}
	way = append(way, root.Val)
	if root.Val == p {
		pa := make([]int, 0)
		for i := range way {
			pa = append(pa, way[i])
		}
		parr = pa
	}
	if root.Val == q {
		qa := make([]int, 0)
		for i := range way {
			qa = append(qa, way[i])
		}
		qarr = qa
	}

	traversal(root.Left, p, q, way)
	traversal(root.Right, p, q, way)
}
