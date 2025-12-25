package linkedList

func IsPalindrome(head *ListNode) bool {
	// num>=1
	if head.Next == nil {
		return true
	}

	midprev := findMiddlePrev(head)
	middle := midprev.Next
	midprev.Next = nil

	newHead := reverseList(middle)
	for head != nil {
		if head.Val != newHead.Val {
			return false
		}
		//链表遍历的时候总忘记写这个
		head = head.Next
		newHead = newHead.Next
	}
	return true
}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	var pre, cur *ListNode
//	cur = head
//	for cur != nil {
//		temp := cur.Next
//		cur.Next = pre
//		pre = cur
//		cur = temp
//	}
//	return pre
//}

func findMiddlePrev(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next.Next == nil {
		return head
	}
	slow := head
	fast := head.Next.Next
	// 少写了 fast != nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
