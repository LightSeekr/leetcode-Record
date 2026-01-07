package tree

/**
1. 直径的定义：路径长度 = 路径上的边数。
2. 关键转化：对于任意一个节点 node，经过该节点的“最长路径”长度等于：左子树的最大深度 + 右子树的最大深度
3. 全局最优：整棵树的直径，就是所有节点中，上述计算结果（左+右）的最大值。

注意：最长路径不一定经过根节点 root，它可能完全包含在某个子树内部。
我们需要一个递归函数（DFS），它要做两件事：
1. 向上汇报：告诉父节点，“以我为起点的最长单侧路径（深度）”是多少。（为了给父节点算它的直径用）。
2. 更新全局最大值：计算“穿过当前节点的最长路径”，并尝试更新全局的 ans。
*/

func diameterOfBinaryTree(root *TreeNode) int {
	/**
	  要想直径最长
	  那么就是两个节点之间路径最长
	  那么一定就在左右两边一边一个。
	  这没办法直接求直径
	  借助求深度来求
	  **/
	maxDiameter := 0
	var getDepth func(root *TreeNode) int
	getDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := getDepth(root.Left)
		rightDepth := getDepth(root.Right)
		maxDiameter = max(maxDiameter, leftDepth+rightDepth)

		return max(leftDepth, rightDepth) + 1
	}

	getDepth(root)
	return maxDiameter

}
