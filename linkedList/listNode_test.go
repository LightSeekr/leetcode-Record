package linkedList_test

import (
	"fmt"
	"go_interview/leetcode/linkedList"
	"testing"
)

func BuildListNode(arr []int) *linkedList.ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &linkedList.ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &linkedList.ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return head
}

func TestPartition(t *testing.T) {
	head := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	head = linkedList.Partition(head, 3)

}

func TestDeleteNode(t *testing.T) {
	head := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	linkedList.DeleteNode(head.Next)
	linkedList.PrintList(head)
}
func TestConstructor(t *testing.T) {
	head := linkedList.BuildRandomList()
	linkedList.PrintList1(head)
	res := linkedList.CopyRandomList(head)
	linkedList.PrintList1(res)
}

func TestMiddleNode(t *testing.T) {
	head := BuildListNode([]int{1, 2, 3, 4, 5})
	res := linkedList.MiddleNode(head)
	t.Log(res.Val)
	head = BuildListNode([]int{1, 2, 3, 4, 5, 6})
	res = linkedList.MiddleNode(head)
	t.Log(res.Val)
}
func TestFindMiddlePrevious(t *testing.T) {
	head := BuildListNode([]int{2, 6, 0, 7, 5, 9})
	res := linkedList.FindMiddlePrevious(head)
	t.Log(res.Val)
	head = BuildListNode([]int{9, 4, 5, 2, 0})
	res = linkedList.FindMiddlePrevious(head)
	t.Log(res.Val)
}
func TestPairSum(t *testing.T) {
	head := BuildListNode([]int{5, 4, 2, 1})
	res := linkedList.PairSum(head)
	t.Log(res)
}

func TestOddEvenList(t *testing.T) {
	head := BuildListNode([]int{1, 2, 3, 4, 5})
	res := linkedList.OddEvenList(head)
	linkedList.PrintList(res)
}

func TestIsPalindrome(t *testing.T) {
	head := BuildListNode([]int{2, 5, 3, 3, 5, 2})
	res := linkedList.IsPalindrome(head)
	t.Log(res)
}

func TestDetectCycle(t *testing.T) {
	head := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	// 加上环
	q := head.Next.Next
	p := head
	for p.Next != nil {
		p = p.Next
	}
	p.Next = q
	fmt.Println("p.val: ", p.Val)
	fmt.Println("q.Val: ", q.Val)
	res := linkedList.DetectCycle(head)
	t.Log(res.Val)
}

func TestMergeTwoLists(t *testing.T) {
	l1 := BuildListNode([]int{1, 2, 3, 4})
	l2 := BuildListNode([]int{5, 6, 7})
	l3 := linkedList.MergeTwoLists(l1, l2)
	linkedList.PrintList(l3)
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := BuildListNode([]int{9, 9, 9, 9, 9, 9})
	l2 := BuildListNode([]int{9, 9, 9})
	l3 := linkedList.AddTwoNumbers(l1, l2)
	linkedList.PrintList(l3)
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	res := linkedList.RemoveNthFromEnd(head, 6)
	linkedList.PrintList(res)

}

func TestSwapPairs(t *testing.T) {
	head := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	res := linkedList.SwapPairs(head)
	linkedList.PrintList(res)
}
