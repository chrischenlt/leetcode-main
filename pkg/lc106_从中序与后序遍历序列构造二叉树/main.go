package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	// 左中右
	// 左右中
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	postOrderLastValue := postorder[len(postorder)-1]
	root := &TreeNode{Val: postOrderLastValue}

	var index int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == postOrderLastValue {
			index = i
			break
		}
	}
	leftInorder := inorder[:index]
	rightInorder := inorder[index+1:]
	root.Left = buildTree(leftInorder, postorder[:len(leftInorder)])
	if len(leftInorder)+1 < len(postorder) {
		root.Right = buildTree(rightInorder, postorder[len(leftInorder):len(postorder)-1])
	}
	return root
}
