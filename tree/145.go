package tree

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	var postorder func(root *TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		res = append(res, root.Val)
	}
	postorder(root)
	return res
}
