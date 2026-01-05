package tree

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkValidBST(root, math.MinInt64, math.MaxInt64)
}

func checkValidBST(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}
	// 严格大于
	if root.Val <= minVal || root.Val >= maxVal {
		return false
	}

	return checkValidBST(root.Left, minVal, root.Val) && checkValidBST(root.Right, root.Val, maxVal)
}

func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	pre := math.MinInt64 // 记录前一个节点的元素
	var inorder func(root *TreeNode) bool
	inorder = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		if !inorder(root.Left) {
			return false
		}
		// 2. 处理当前节点
		// 如果当前节点小于等于前一个节点，说明不是严格递增，返回 false
		if root.Val <= pre {
			return false
		}
		// 更新 pre
		pre = root.Val

		return inorder(root.Right)
	}

	return inorder(root)
}
