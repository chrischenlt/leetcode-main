package lc513_找树左下角的值

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代
func findBottomLeftValue(root *TreeNode) int {
	list := []*TreeNode{root}
	var res int
	for len(list) > 0 {
		length := len(list)
		newList := make([]*TreeNode, 0)
		res = list[0].Val
		for i := 0; i < length; i++ {
			if list[i].Left != nil {
				newList = append(newList, list[i].Left)
			}
			if list[i].Right != nil {
				newList = append(newList, list[i].Right)
			}
		}
		list = newList
	}
	return res
}

// 递归
var res, maxDepth int

func findBottomLeftValue1(root *TreeNode) int {
	res, maxDepth = 0, 0
	a(root, 1)
	return res
}
func a(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if depth > maxDepth {
			maxDepth = depth
			res = root.Val
		}
		return
	}
	a(root.Left, depth+1)
	a(root.Right, depth+1)
}
