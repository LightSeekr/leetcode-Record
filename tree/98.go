package tree

import "math"

func isValidBST(root *TreeNode) bool {
	// 初始范围是无穷小到无穷大
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	// 1. 空节点是合法的
	if root == nil {
		return true
	}

	// 2. 判断当前节点是否越界
	// 必须是 严格小于 和 严格大于，所以用 <= 和 >=
	if root.Val <= lower || root.Val >= upper {
		return false
	}

	// 3. 递归检查子节点
	// 往左走：上界变成 root.Val
	// 往右走：下界变成 root.Val
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

func isValidBST2(root *TreeNode) bool {
	// 用一个变量记录前一个节点的值，初始为极小值
	pre := math.MinInt64

	// 定义 DFS 函数
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		// 1. 先遍历左子树
		if !dfs(node.Left) {
			return false
		}

		// 2. 处理当前节点
		// 如果当前节点小于等于前一个节点，说明不是严格递增，返回 false
		if node.Val <= pre {
			return false
		}
		// 更新 pre
		pre = node.Val

		// 3. 后遍历右子树
		return dfs(node.Right)
	}

	return dfs(root)
}
