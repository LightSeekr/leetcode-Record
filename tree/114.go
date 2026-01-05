package tree

/*
*给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，
其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/
func flatten(root *TreeNode) {
	list := make([]*TreeNode, 0)
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		list = append(list, root)
		preOrder(root.Left)
		preOrder(root.Right)
	}
	// 前序遍历
	preOrder(root)

	for i := 0; i < len(list)-1; i++ {
		list[i].Left, list[i].Right = nil, list[i+1]
	}
	if len(list) > 0 {
		list[len(list)-1].Left = nil
		list[len(list)-1].Right = nil
	}
}
