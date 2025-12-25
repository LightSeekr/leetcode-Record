package linkedList

func DeleteNode(node *ListNode) {
	cur := node
	for cur.Next.Next != nil {
		cur.Val = cur.Next.Val
		cur = cur.Next
	}
	cur.Val = cur.Next.Val
	cur.Next = nil
}
