package LinkedList

func DeleteMiddle(head *ListNode) *ListNode {
	// 边界 case 特判
	if head == nil || head.Next == nil {
		return nil
	}
	return head

}
