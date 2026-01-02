package tree

func kthSmallest(root *TreeNode, k int) int {
	var result int

	// 定义中序遍历函数
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		// 如果节点为空，或者已经找到了答案（k==0），直接返回
		if node == nil || k == 0 {
			return
		}

		// 1. 先去左边
		inorder(node.Left)

		// 2. 处理当前节点
		k-- // 每访问一个节点，k 就减 1
		if k == 0 {
			result = node.Val
			return
		}

		// 3. 再去右边
		inorder(node.Right)
	}

	inorder(root)
	return result
}
