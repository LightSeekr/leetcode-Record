package tree

func flatten(root *TreeNode) {
	var queue []*TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		queue = append(queue, node)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	for i := 0; i < len(queue)-1; i++ {
		queue[i].Left, queue[i].Right = nil, queue[i+1]
	}
	if len(queue) > 0 {
		queue[len(queue)-1].Left = nil
		queue[len(queue)-1].Right = nil
	}

}
