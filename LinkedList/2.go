package LinkedList

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	num := 0
	p, q := l1, l2
	dummy := &ListNode{}
	tail := dummy
	for p != nil || q != nil {
		if p != nil {
			num += p.Val
			p = p.Next
		}
		if q != nil {
			num += q.Val
			q = q.Next
		}
		tail.Next = &ListNode{Val: num % 10, Next: nil}
		tail = tail.Next
		num /= 10
	}
	if num != 0 {
		tail.Next = &ListNode{Val: num}
		tail = tail.Next
	}

	return dummy.Next
}
