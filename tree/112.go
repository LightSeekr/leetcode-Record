package tree

func hasPathSum(root *TreeNode, targetSum int) bool {
	// 1. 递归终止条件：空节点
	// 注意：即使 targetSum 为 0，空树也不包含路径，所以返回 false
	if root == nil {
		return false
	}

	// 2. 判断是否是叶子节点
	if root.Left == nil && root.Right == nil {
		// 如果是叶子节点，判断当前值是否等于剩下的目标值
		return root.Val == targetSum
	}

	// 3. 递归处理左右子树
	// 只要有一边能找到路径即可，所以用 || (或)
	// 传入新的 targetSum = targetSum - root.Val
	return hasPathSum(root.Left, targetSum-root.Val) ||
		hasPathSum(root.Right, targetSum-root.Val)
}

func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// 两个队列，一个存节点，一个存到达该节点时的路径和
	nodeQueue := []*TreeNode{root}
	valQueue := []int{root.Val}

	for len(nodeQueue) > 0 {
		// 出队
		currNode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		currVal := valQueue[0]
		valQueue = valQueue[1:]

		// 判断是否是叶子节点
		if currNode.Left == nil && currNode.Right == nil {
			if currVal == targetSum {
				return true
			}
			continue
		}

		// 左子节点入队
		if currNode.Left != nil {
			nodeQueue = append(nodeQueue, currNode.Left)
			valQueue = append(valQueue, currVal+currNode.Left.Val)
		}

		// 右子节点入队
		if currNode.Right != nil {
			nodeQueue = append(nodeQueue, currNode.Right)
			valQueue = append(valQueue, currVal+currNode.Right.Val)
		}
	}

	return false
}
