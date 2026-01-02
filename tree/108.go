package tree

func sortedArrayToBST(nums []int) *TreeNode {
	var build func(nums []int, left, right int) *TreeNode

	build = func(nums []int, left, right int) *TreeNode {
		// 递归终止条件：左边界超过右边界，说明没有元素了
		if left > right {
			return nil
		}

		// 1. 找到中间位置
		// (left + right) / 2 在偶数个元素时会选择左偏中间的那个，这也是符合题意的
		mid := (left + right) / 2

		// 2. 创建当前根节点
		root := &TreeNode{Val: nums[mid]}

		// 3. 递归构建左子树 (范围: left 到 mid-1)
		root.Left = build(nums, left, mid-1)

		// 4. 递归构建右子树 (范围: mid+1 到 right)
		root.Right = build(nums, mid+1, right)

		return root
	}

	// 调用辅助递归函数
	return build(nums, 0, len(nums)-1)
}

func sortedArrayToBST2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := (len(nums) - 1) / 2

	root := &TreeNode{Val: nums[mid]}

	root.Left = sortedArrayToBST2(nums[0 : mid-1])
	root.Right = sortedArrayToBST2(nums[mid+1 : len(nums)-1])
	return root
}
