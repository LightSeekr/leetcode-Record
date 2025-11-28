package LinkedList

func SwapPairs(head *ListNode) *ListNode {
	// 边界 case
	return swapPairs(head)
}

// 递归好写，迭代难想 吃饭去，回来再继续写

func swapPairs(head *ListNode) *ListNode {
	//// 递归写法
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//swapHead := swapPairs(head.Next.Next)
	//p, q := head, head.Next
	//q.Next = p
	//p.Next = swapHead
	//return q

	// 迭代写法
	dummy := &ListNode{}
	dummy.Next = head
	node0 := dummy
	node1 := dummy.Next

	for node1 != nil && node1.Next != nil {
		node2 := node1.Next
		node3 := node2.Next

		node0.Next = node2
		node2.Next = node1
		node1.Next = node3

		//
		node0 = node1
		node1 = node3
	}
	return dummy.Next

}
