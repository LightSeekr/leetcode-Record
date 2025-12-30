package linkedList

// MergeTwoLists 这算一道母题啊，不知道记录了没有
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	p, q := list1, list2
	dummy := &ListNode{}
	tail := dummy
	for p != nil && q != nil {
		if p.Val < q.Val {
			tail.Next = p
			p = p.Next
		} else {
			tail.Next = q
			q = q.Next
		}
		tail = tail.Next
	}

	if p != nil {
		tail.Next = p
	}
	if q != nil {
		tail.Next = q
	}
	return dummy.Next
}
