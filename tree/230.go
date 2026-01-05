package tree

/**
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（k 从 1 开始计数）。
*/

// 直接一个中序遍历
func kthSmallest(root *TreeNode, k int) int {
	res := make([]int, 0)
	var inOrder func(root *TreeNode)

	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		res = append(res, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)
	return res[k-1]
}

// 对中序遍历做一些优化
func kthSmallest2(root *TreeNode, k int) int {
	res := 0
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil || k == 0 {
			return
		}

		inOrder(root.Left)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		inOrder(root.Right)
	}

	inOrder(root)
	return res
}
