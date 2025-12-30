package linkedList

import (
	"fmt"
	"testing"
)

func BuildListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return head
}

func TestPartition(t *testing.T) {
	head := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	head = Partition(head, 3)

}

func TestDeleteNode(t *testing.T) {
	head := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	DeleteNode(head.Next)
	PrintList(head)
}
func TestConstructor(t *testing.T) {
	head := BuildRandomList()
	PrintList1(head)
	res := CopyRandomList(head)
	PrintList1(res)
}

func TestMiddleNode(t *testing.T) {
	head := BuildListNode([]int{1, 2, 3, 4, 5})
	res := MiddleNode(head)
	t.Log(res.Val)
	head = BuildListNode([]int{1, 2, 3, 4, 5, 6})
	res = MiddleNode(head)
	t.Log(res.Val)
}
func TestFindMiddlePrevious(t *testing.T) {
	head := BuildListNode([]int{2, 6, 0, 7, 5, 9})
	res := FindMiddlePrevious(head)
	t.Log(res.Val)
	head = BuildListNode([]int{9, 4, 5, 2, 0})
	res = FindMiddlePrevious(head)
	t.Log(res.Val)
}
func TestPairSum(t *testing.T) {
	head := BuildListNode([]int{5, 4, 2, 1})
	res := PairSum(head)
	t.Log(res)
}

func TestOddEvenList(t *testing.T) {
	head := BuildListNode([]int{1, 2, 3, 4, 5})
	res := OddEvenList(head)
	PrintList(res)
}

func TestIsPalindrome(t *testing.T) {
	head := BuildListNode([]int{2, 5, 3, 3, 5, 2})
	res := IsPalindrome(head)
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
	res := DetectCycle(head)
	t.Log(res.Val)
}

func TestMergeTwoLists(t *testing.T) {
	l1 := BuildListNode([]int{1, 2, 3, 4})
	l2 := BuildListNode([]int{5, 6, 7})
	l3 := MergeTwoLists(l1, l2)
	PrintList(l3)
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := BuildListNode([]int{9, 9, 9, 9, 9, 9})
	l2 := BuildListNode([]int{9, 9, 9})
	l3 := AddTwoNumbers(l1, l2)
	PrintList(l3)
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	res := RemoveNthFromEnd(head, 6)
	PrintList(res)
}

func TestSwapPairs(t *testing.T) {
	head := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	res := SwapPairs(head)
	PrintList(res)
}

func TestReverseBetween(t *testing.T) {
	head := BuildListNode([]int{1, 2, 3, 4, 5})
	res := reverseBetween(head, 2, 4)
	PrintList(res)
}

func TestReverseKGroup(t *testing.T) {
	head := BuildListNode([]int{1, 2, 3, 4, 5})

	res := reverseKGroup(head, 2)
	PrintList(res)

}
