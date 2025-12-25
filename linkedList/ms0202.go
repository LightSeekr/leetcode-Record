package linkedList

// 母题
func KthToLast(head *ListNode, k int) int {
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
	return cur.Val
}
