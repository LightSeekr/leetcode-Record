package linkedList

/*
*
遍历链表，找到 left 位置的前一个节点，我们称之为 pre。
在反转过程中，pre 始终指向反转区域的前一个位置，保持不动。

cur 为反转区域的第一个节点（即 pre.Next）。
在 left 到 right 的区间内，我们需要不断地把 cur 后面的那个节点（next）移到 pre 的后面。
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy

	// 让 pre 指向翻转区域的前一个节点。
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	// 翻转的第一个节点
	cur := pre.Next

	for i := 1; i < right-left+1; i++ {
		nxt := cur.Next

		// cur 连接到 next 的下一个节点 (跳过 next)
		cur.Next = nxt.Next

		// 将节点插入到 pre 的后面
		nxt.Next = pre.Next
		pre.Next = nxt
	}
	return dummy.Next
}
