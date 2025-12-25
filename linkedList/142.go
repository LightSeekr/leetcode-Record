package linkedList

func DetectCycle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast

		}
	}
	return nil
}
