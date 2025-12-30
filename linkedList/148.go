package linkedList

/*
*

	写过一个归并排序，递归版本的。
	先找中点，把链表断成两截（这个有AI写的一个东西）
	sortList(左)
	sortList(右)
	merge 两个链表，
	链表的归并排序

给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	middle := FindMiddlePrevious(head)
	newHead := middle.Next
	middle.Next = nil
	s1 := sortList(head)
	s2 := sortList(newHead)
	return MergeTwoLists(s1, s2)
}
