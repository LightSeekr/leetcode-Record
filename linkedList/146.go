package linkedList

type LNode struct {
	Front *LNode
	Next  *LNode
	Key   int
	Val   int
}

type DList struct {
	Head *LNode
	Tail *LNode
	Len  int
}

type LRUCache struct {
	capacity int
	DList
	M1 map[int]*LNode
}

func Constructor(capacity int) LRUCache {
	l1 := LRUCache{
		capacity: capacity,
		DList: DList{
			Head: &LNode{
				Val: -1,
			},
			Tail: &LNode{
				Val: -2,
			},
			Len: 0,
		},
		M1: make(map[int]*LNode, capacity),
	}
	l1.Head.Next = l1.Tail
	l1.Tail.Front = l1.Head
	return l1
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.M1[key]; !ok {
		return -1
	} else {
		// 找到了 需要将这个节点，放到链表头部
		l.Refresh(node)
		return node.Val
	}
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.M1[key]; ok {
		// 存在 则更新缓存
		node.Val = value
		l.Refresh(node)
	} else {
		// 头插链表
		n := &LNode{
			Key: key,
			Val: value,
		}
		p1 := l.Head
		p2 := p1.Next

		n.Next = p1.Next
		p1.Next = n

		n.Front = p2.Front
		p2.Front = n

		l.Len++
		//如果 len>cap 删除尾部节点 // map 删除节点
		if l.Len > l.capacity {
			l.Delete()
			l.Len--
		}
		// 确定后，再插入哈希表
		l.M1[key] = n
	}

}

func (l *LRUCache) Delete() {
	dNode := l.Tail.Front
	// 从哈希表中删除
	delete(l.M1, dNode.Key)

	// 从双向链表中删除
	p := dNode.Front
	p.Next = l.Tail
	l.Tail.Front = p
	dNode.Front = nil
	dNode.Next = nil
}

func (l *LRUCache) Refresh(node *LNode) {

	f := node.Front
	n := node.Next

	f.Next = n
	n.Front = f

	node.Next = nil
	node.Front = nil

	node.Next = l.Head.Next
	node.Front = l.Head
	l.Head.Next = node
	node.Next.Front = node
}
