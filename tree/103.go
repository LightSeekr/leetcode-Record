package tree

import (
	"slices"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	cnt := 0
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		size := len(queue)
		cnt += 1
		ans := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			ans = append(ans, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		//fmt.Printf("ans:%+v\n",ans)
		// 偶数
		if cnt%2 == 0 {
			slices.Reverse(ans)
		}
		//fmt.Printf("ans:%+v\n",ans)
		res = append(res, ans)
	}
	return res
}
