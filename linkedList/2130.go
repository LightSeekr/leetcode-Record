package linkedList

import "fmt"

func PairSum(head *ListNode) int {
	// 节点数>=2 且为偶数
	if head.Next.Next == nil {
		return head.Val + head.Next.Val
	}
	// 找到 pre
	pre := FindMiddlePrevious(head)
	// 断开
	middle := pre.Next
	pre.Next = nil
	PrintList(head)
	fmt.Println("======================================================")
	PrintList(middle)

	// 反转 middle
	newHead := reverseList(middle)
	fmt.Println("======================================================")
	PrintList(newHead)
	fmt.Println("======================================================")
	// 逐个相加 取最大
	p, q := head, newHead
	res := 0
	for p != nil && q != nil {
		temp := p.Val + q.Val
		if temp > res {
			res = temp
		}
		q = q.Next
		p = p.Next
	}
	return res
}

//
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	newHead := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return newHead
//
//}

//func FindMiddlePrevious(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return nil
//	}
//	// 处理只有两个节点的特殊情况
//	if head.Next.Next == nil {
//		return head
//	}
//	slow, fast := head, head.Next.Next
//	for fast != nil && fast.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	return slow
//}
