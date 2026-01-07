package tree

/*
*
给定一个二叉树，判断它是否是 平衡二叉树
*
*/

func isBalanced(root *TreeNode) bool {
	_, balanced := checkBalanced(root)
	return balanced
}

func checkBalanced(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftH, leftOk := checkBalanced(node.Left)
	if !leftOk {
		return 0, false
	} // 提前退出

	rightH, rightOk := checkBalanced(node.Right)
	if !rightOk {
		return 0, false
	} // 提前退出

	// 检查差值
	diff := leftH - rightH
	if diff < -1 || diff > 1 {
		return 0, false
	}

	// 返回 max(left, right) + 1 和 true
	return max(leftH, rightH) + 1, true
}

func isBalanced2(root *TreeNode) bool {
	// 左子树平衡 && 右子树平衡 && 左右子树高度差不能超过 1
	// 把求高度 和 算平衡这两件事统一起来
	if root == nil {
		return true
	}
	leftDepth := getTreeDepth(root.Left)

	if leftDepth == -1 {
		return false
	}
	rightDepth := getTreeDepth(root.Right)
	if rightDepth == -1 {
		return false
	}
	return abs(leftDepth-rightDepth) <= 1

}

func getTreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := getTreeDepth(root.Left)
	if leftDepth == -1 {
		return -1
	}

	rightDepth := getTreeDepth(root.Right)
	if rightDepth == -1 {
		return -1
	}

	if abs(leftDepth-rightDepth) > 1 {
		return -1
	}
	return max(leftDepth, rightDepth) + 1
}
func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return a * -1
	}
}
