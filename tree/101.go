package tree

/**
两棵树（记为 A 和 B）互为镜像的条件是：
根节点值相等：A.Val == B.Val。
交叉匹配：
A 的左孩子 == B 的右孩子。
A 的右孩子 == B 的左孩子。
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 调用辅助函数，比较左子树和右子树
	return check(root.Left, root.Right)
}

// 辅助函数：判断两个节点 p 和 q 是否互为镜像
func check(p, q *TreeNode) bool {
	// 1. 两个都为空，是对称的
	if p == nil && q == nil {
		return true
	}

	// 2. 一个为空一个不为空，或者值不相等，不是对称的
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	// 3. 递归检查内部子节点：
	// p 的左边 vs q 的右边
	// p 的右边 vs q 的左边
	return check(p.Left, q.Right) && check(p.Right, q.Left)
}
