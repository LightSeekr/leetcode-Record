package linkedList

func reverseList(head *ListNode) *ListNode {
	/**
	这样写可以，但是 最后加一个 head.Next = nil 想当于是特殊处理了，不太好。
	*/
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//prev := head
	//cur := head.Next
	//for cur != nil {
	//	next := cur.Next
	//	cur.Next = prev
	//	prev = cur
	//	cur = next
	//}
	//head.Next = nil
	//return prev
	// 始终坚持， pre 指向翻转区域的前一个节点。下一题也是这个解法
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur *ListNode
	cur = head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
