package linkedList

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	// 边界 case
	if head == nil {
		return head
	}

	p, q := head, head.Next
	// 把原链表的每一个节点后都加上新节点,
	// 赋值新节点的 next
	for p != nil {
		q = p.Next
		temp := &Node{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}
		temp.Next = q
		p.Next = temp
		p = q
	}

	// 新节点中的每一个节点的 random 赋予值
	p = head
	for p != nil {
		q = p.Next
		if p.Random != nil {
			q.Random = p.Random.Next //
		} else {
			q.Random = nil
		}
		p = q.Next
	}
	PrintList1(head)

	fakeHead := &Node{
		Next:   nil,
		Random: nil,
	}
	p, q = head, fakeHead
	for p != nil {
		temp := p.Next
		p.Next = temp.Next
		q.Next = temp
		q = q.Next
		p = p.Next
	}
	return fakeHead.Next
}

// [[7,null],[13,0],[11,4],[10,2],[1,0]]
func BuildRandomList() *Node {

	var Node0, Node1, Node2, Node3, Node4 *Node
	Node0 = &Node{
		Val:    7,
		Next:   nil,
		Random: nil,
	}

	Node1 = &Node{
		Val:    13,
		Next:   nil,
		Random: Node0,
	}
	Node0.Next = Node1

	Node2 = &Node{
		Val:    11,
		Next:   nil,
		Random: nil,
	}
	Node1.Next = Node2

	Node3 = &Node{
		Val:    10,
		Next:   nil,
		Random: Node2,
	}
	Node2.Next = Node3

	Node4 = &Node{
		Val:    1,
		Next:   nil,
		Random: Node0,
	}
	Node3.Next = Node4
	Node2.Random = Node4
	return Node0
}

func PrintList1(head *Node) {
	for head != nil {
		fmt.Printf("%p ", head)
		head = head.Next
	}
	fmt.Printf("\n======================================\n")
}

func PrintNode(head *Node) {
	q := head
	//for head != nil {
	//	fmt.Printf("%p ", head.Val)
	//	head = head.Next
	//}
	//fmt.Printf("\n======================================\n")

	head = q
	for head != nil {
		if head.Random != nil {
			fmt.Printf("%d ", head.Random.Val)
		} else {
			fmt.Print("nil ")
		}
		head = head.Next
	}
	fmt.Printf("\n======================================\n")
}
