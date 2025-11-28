package LinkedList_test

import (
	"fmt"
	"go_interview/leetcode/LinkedList"
	"testing"
)

func BuildListNode(arr []int) *LinkedList.ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &LinkedList.ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &LinkedList.ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return head
}

func TestPartition(t *testing.T) {
	head := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	head = LinkedList.Partition(head, 3)

}

func TestDeleteNode(t *testing.T) {
	head := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	LinkedList.DeleteNode(head.Next)
	LinkedList.PrintList(head)
}
func TestConstructor(t *testing.T) {
	head := LinkedList.BuildRandomList()
	LinkedList.PrintList1(head)
	res := LinkedList.CopyRandomList(head)
	LinkedList.PrintList1(res)
}

func TestMiddleNode(t *testing.T) {
	head := BuildListNode([]int{1, 2, 3, 4, 5})
	res := LinkedList.MiddleNode(head)
	t.Log(res.Val)
	head = BuildListNode([]int{1, 2, 3, 4, 5, 6})
	res = LinkedList.MiddleNode(head)
	t.Log(res.Val)
}
func TestFindMiddlePrevious(t *testing.T) {
	head := BuildListNode([]int{2, 6, 0, 7, 5, 9})
	res := LinkedList.FindMiddlePrevious(head)
	t.Log(res.Val)
	head = BuildListNode([]int{9, 4, 5, 2, 0})
	res = LinkedList.FindMiddlePrevious(head)
	t.Log(res.Val)
}
func TestPairSum(t *testing.T) {
	head := BuildListNode([]int{5, 4, 2, 1})
	res := LinkedList.PairSum(head)
	t.Log(res)
}

func TestOddEvenList(t *testing.T) {
	head := BuildListNode([]int{1, 2, 3, 4, 5})
	res := LinkedList.OddEvenList(head)
	LinkedList.PrintList(res)
}

func TestIsPalindrome(t *testing.T) {
	head := BuildListNode([]int{2, 5, 3, 3, 5, 2})
	res := LinkedList.IsPalindrome(head)
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
	res := LinkedList.DetectCycle(head)
	t.Log(res.Val)
}

func TestMergeTwoLists(t *testing.T) {
	l1 := BuildListNode([]int{1, 2, 3, 4})
	l2 := BuildListNode([]int{5, 6, 7})
	l3 := LinkedList.MergeTwoLists(l1, l2)
	LinkedList.PrintList(l3)
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := BuildListNode([]int{9, 9, 9, 9, 9, 9})
	l2 := BuildListNode([]int{9, 9, 9})
	l3 := LinkedList.AddTwoNumbers(l1, l2)
	LinkedList.PrintList(l3)
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	res := LinkedList.RemoveNthFromEnd(head, 6)
	LinkedList.PrintList(res)

}

func TestSwapPairs(t *testing.T) {
	head := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	res := LinkedList.SwapPairs(head)
	LinkedList.PrintList(res)
}
