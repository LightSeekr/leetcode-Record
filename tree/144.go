package tree

func preorderTraversal(root *TreeNode) []int {
	var res []int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// Preorder: visit node, then left subtree, then right subtree.
		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return res
}
