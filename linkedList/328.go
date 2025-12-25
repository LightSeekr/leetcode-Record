package linkedList

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 奇数节点头
	oddHead := &ListNode{Next: nil}
	p := oddHead

	// 偶数节点头
	evenHead := &ListNode{Next: nil}
	q := evenHead

	cnt := 1
	for head != nil {
		if cnt%2 == 0 {
			// 偶
			q.Next = head
			head = head.Next
			q = q.Next
			q.Next = nil
		} else {

			p.Next = head
			head = head.Next
			p = p.Next
			p.Next = nil
		}
		cnt++
	}

	p.Next = evenHead.Next
	return oddHead.Next

}
