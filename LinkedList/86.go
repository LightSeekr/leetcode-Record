package LinkedList

func Partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fakeHead1, fakeHead2 := &ListNode{Next: nil}, &ListNode{Next: nil}
	f, p, q := head, fakeHead1, fakeHead2
	for f != nil {
		if f.Val < x {
			p.Next = f
			p = p.Next
			f = f.Next
			p.Next = nil
		} else {
			q.Next = f
			q = q.Next
			f = f.Next
			q.Next = nil
		}
	}
	p.Next = fakeHead2.Next
	PrintList(fakeHead1.Next)
	return fakeHead1.Next

}
