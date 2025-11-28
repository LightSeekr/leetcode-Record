package LinkedList

// MiddleNode
func MiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// FindMiddlePrevious 返回链表中间节点的前一个节点
// 对于奇数长度链表 [2, 6, 0, 7, 5]，返回值为6的节点（中间节点0的前一个）
// 对于偶数长度链表 [4, 3, 2, 1]，返回值为3的节点（第一个中间节点）
func FindMiddlePrevious(head *ListNode) *ListNode {
	// 处理特殊情况：空链表或只有一个节点
	if head == nil || head.Next == nil {
		return nil // 没有前一个节点
	}
	// 特殊情况：只有两个节点的链表
	if head.Next.Next == nil {
		return head // 返回第一个节点
	}

	// 初始化慢指针和快指针
	slow := head
	fast := head.Next.Next

	// 遍历链表，当快指针到达末尾时，慢指针正好在中间节点的前一个位置
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
