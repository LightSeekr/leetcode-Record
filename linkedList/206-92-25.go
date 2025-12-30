package linkedList

/**
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
*
我们可以把整个链表分为三个部分来看待每一次的操作：
1. 已处理区域：dummy 到 pre。
2. 待反转区域（长度为 k）：start 到 end。
3. 未处理区域：next 及其后面。
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	dummy := &ListNode{Val: -1, Next: head}
	pre := dummy
	end := pre
	for end.Next != nil {
		start := pre.Next // 待反转区域的第一个节点

		// 1. 尝试向后移动 k 步，看是否有足够的节点
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		// 如果 end 为 nil，说明剩余节点不足 k 个，不需要反转，直接跳出
		if end == nil {
			break
		}

		// 2. 记录关键节点

		nextGroup := end.Next // 下一组的起始节点

		// 3. 断开链表，方便独立的 reverse 函数处理
		end.Next = nil
		// 4. 反转当前这 k 个节点
		newNode := reverse(start)
		// pre.Next 指向反转后的新头节点
		pre.Next = newNode

		// 5. 重新连接
		// 反转后，start 变成了尾节点，将其指向下一组
		start.Next = nextGroup
		// 6. 重置指针，准备下一轮
		pre = start
		end = pre
	}

	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
