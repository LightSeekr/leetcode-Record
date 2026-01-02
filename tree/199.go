package tree

func rightSideView(root *TreeNode) []int {
	ns := levelOrder(root)
	res := make([]int, 0)
	for _, n := range ns {
		res = append(res, n[len(n)-1])
	}
	return res
}
