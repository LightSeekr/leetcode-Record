package tree

/*
*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

*
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 基础情况：如果节点为空，或者找到了 p 或 q，直接返回
	if root == nil || root == p || root == q {
		return root
	}

	// 递归查找左子树和右子树
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果左右子树都找到了节点，说明 p 和 q 分别在左右子树，当前节点就是最近公共祖先
	if left != nil && right != nil {
		return root
	}

	// 如果只有左子树找到，返回左子树的结果
	if left != nil {
		return left
	}

	// 如果只有右子树找到，返回右子树的结果
	return right
}
