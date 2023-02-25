package lc101_对称二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	list := []*TreeNode{root.Left, root.Right}
	for len(list) > 1 {
		left, right := 0, len(list)-1
		newList := make([]*TreeNode, 0)
		for left < right {
			a1, a2 := list[left], list[right]
			left++
			right--
			if a1 == nil && a2 == nil {
				continue
			}
			if a1 == nil || a2 == nil || a1.Val != a2.Val {
				return false
			}
			newList = append(newList, a1.Left)
			newList = append(newList, a1.Right)
			newList = append(newList, a2.Left)
			newList = append(newList, a2.Right)
		}
		list = newList
	}
	if len(list) > 0 {
		return false
	}

	return true
}

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSymmetricxx(root.Left, root.Right)
}

func isSymmetricxx(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return isSymmetricxx(t1.Left, t2.Right) && isSymmetricxx(t1.Right, t2.Left)
}
