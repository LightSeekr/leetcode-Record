package tree

/**
1 <= preorder.length <= 3000
inorder.length == preorder.length
*/

/*
*
func buildTree(preorder []int, inorder []int) *TreeNode {

		// 1. 根节点永远是 preorder 的第一个元素
		rootVal := preorder[0]
		root := &TreeNode{Val: rootVal}

		// 2. 在 inorder 中找到根节点的位置（索引 i）
		// 这一步替代了 Map，通过遍历查找
		i := 0
		for ; i < len(inorder); i++ {
			if inorder[i] == rootVal {
				break
			}
		}

		// 此时，i 不仅是根节点在 inorder 中的索引
		// 也代表了【左子树的节点数量】

		// 3. 切分切片并递归
		// 左子树：
		//   preorder: 从下标 1 开始，长度为 i 的区间 -> preorder[1 : 1+i]
		//   inorder:  根节点左边的所有元素 -> inorder[0 : i]
		root.Left = buildTree(preorder[1:1+i], inorder[:i])

		// 右子树：
		//   preorder: 剩下的部分 -> preorder[1+i : ]
		//   inorder:  根节点右边的所有元素 -> inorder[i+1 : ]
		root.Right = buildTree(preorder[1+i:], inorder[i+1:])

		return root
	}
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}

	k := -1
	for i := 0; i < len(inorder); i++ {
		if root.Val == inorder[i] {
			k = i
			break
		}
	}

	root.Left = buildTree(preorder[1:1+k], inorder[:k])
	root.Right = buildTree(preorder[k+1:], inorder[k+1:])

	return root
}

// 使用 Map 而不是 数组
func buildTree2(preorder []int, inorder []int) *TreeNode {
	// 1. 构建哈希表，快速定位 root 在 inorder 中的位置
	indexMap := make(map[int]int)
	for i, v := range inorder {
		indexMap[v] = i
	}

	// preorderIndex 用来记录当前我们在 preorder 数组中用到哪了
	// 这是一个闭包变量，随着递归不断前进
	preorderIndex := 0

	// 定义递归函数
	// left, right 表示当前子树在 inorder 数组中的索引范围
	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		// 递归终止条件：范围无效
		if left > right {
			return nil
		}

		// 1. 获取当前根节点的值 (从 preorder 中按顺序拿)
		rootVal := preorder[preorderIndex]
		preorderIndex++ // 指针后移

		// 2. 创建根节点
		root := &TreeNode{Val: rootVal}

		// 3. 在 inorder 中找到根节点的位置
		// 因为题目保证没有重复元素，所以可以直接用 map
		inorderIndex := indexMap[rootVal]

		// 4. 递归构建左右子树
		// 关键点：preorder 的顺序是 根 -> 左 -> 右
		// 所以我们必须先构建左子树，再构建右子树
		// 左子树的范围：当前左边界 到 分割点-1
		root.Left = build(left, inorderIndex-1)

		// 右子树的范围：分割点+1 到 当前右边界
		root.Right = build(inorderIndex+1, right)

		return root
	}

	return build(0, len(inorder)-1)
}
