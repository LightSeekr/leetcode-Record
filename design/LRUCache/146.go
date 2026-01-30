package LRUCache

/**
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/
// 如果要带过期时间的，应该怎么优化，这是个问题

type Node struct {
	Prev *Node
	Next *Node
	key  int
	val  int
}

type LRUCache struct {
	Head *Node
	Tail *Node
	cap  int
	len  int
	Map  map[int]*Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{}
	l.Head = &Node{
		val: 0,
	}
	l.Tail = &Node{
		val: 0,
	}
	// 构建双向链表
	l.Head.Next = l.Tail
	l.Tail.Prev = l.Head

	l.cap = capacity
	l.Map = make(map[int]*Node, capacity)
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Map[key]; ok {
		// 将这个节点放到链表头部去
		this.refresh(node)
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	// 检查 key 是否已存在
	if node, ok := this.Map[key]; ok {
		// key 已存在，更新值并刷新位置
		node.val = value
		this.refresh(node)
		return
	}

	// key 不存在，创建新节点
	node := &Node{
		key: key,
		val: value,
	}
	this.Map[key] = node
	this.len++
	// 将这个节点放到双向链表的头部
	this.refresh(node)

	if this.len > this.cap {
		// 删除尾节点
		// 在 map 中删除
		this.deleteTailNode()
		this.len--
	}
}

func (this *LRUCache) deleteTailNode() {
	node := this.Tail.Prev
	//在 map 中删除
	delete(this.Map, node.key)

	// 在链表中删除
	pre, nxt := node.Prev, node.Next
	pre.Next = nxt
	nxt.Prev = pre
}

func (this *LRUCache) refresh(node *Node) {
	// 如果节点已在链表中，先从原来的位置摘下来
	if node.Prev != nil && node.Next != nil {
		pre := node.Prev
		nxt := node.Next
		pre.Next = nxt
		nxt.Prev = pre
	}

	//放到头部去
	node.Next = this.Head.Next
	node.Prev = this.Head

	this.Head.Next.Prev = node
	this.Head.Next = node
}
