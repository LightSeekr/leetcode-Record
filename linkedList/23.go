package linkedList

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	// 边界条件：只剩一个链表时，直接返回
	if left == right {
		return lists[left]
	}

	// 找到中点
	mid := left + (right-left)/2

	// 递归合并左半部分
	l1 := merge(lists, left, mid)
	// 递归合并右半部分
	l2 := merge(lists, mid+1, right)

	// 合并两个有序链表
	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return dummy.Next
}
