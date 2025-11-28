package LinkedList

// 引入 dummy 可以避免对头节点的特殊判断
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	//sz>=1 head!=nil
	dummyHead := &ListNode{}
	dummyHead.Next = head
	node := FindNthFromEnd(dummyHead, n+1)
	node.Next = node.Next.Next
	return dummyHead.Next
}

func FindNthFromEnd(head *ListNode, k int) *ListNode {
	// head != nil
	pre := head
	for i := 0; i < k; i++ {
		pre = pre.Next
	}
	cur := head
	for pre != nil {
		pre = pre.Next
		cur = cur.Next
	}
	return cur
}
