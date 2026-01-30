package LRUCache

import (
	"sync"
	"time"
)

/**
带超时时间的 LRU 缓存实现（混合方案）
- 被动过期：Get 时检查是否过期
- 主动清理：定期批量清理过期数据
*/

type NodeWithTTL struct {
	Prev     *NodeWithTTL
	Next     *NodeWithTTL
	key      int
	val      int
	expireAt time.Time // 过期时间点
}

type LRUCacheWithTTL struct {
	Head     *NodeWithTTL
	Tail     *NodeWithTTL
	cap      int
	len      int
	Map      map[int]*NodeWithTTL
	mu       sync.RWMutex // 保护并发访问
	stopChan chan struct{}
	ticker   *time.Ticker
}

func ConstructorWithTTL(capacity int, cleanInterval time.Duration) *LRUCacheWithTTL {
	l := &LRUCacheWithTTL{}
	l.Head = &NodeWithTTL{}
	l.Tail = &NodeWithTTL{}

	// 构建双向链表
	l.Head.Next = l.Tail
	l.Tail.Prev = l.Head

	l.cap = capacity
	l.Map = make(map[int]*NodeWithTTL, capacity)
	l.stopChan = make(chan struct{})
	l.ticker = time.NewTicker(cleanInterval)

	// 启动后台定期清理 goroutine
	go l.cleanExpiredPeriodically()

	return l
}

// Get 获取值，如果过期则返回 -1
func (this *LRUCacheWithTTL) Get(key int) int {
	this.mu.Lock()
	defer this.mu.Unlock()

	if node, ok := this.Map[key]; ok {
		// 被动过期检查
		if time.Now().After(node.expireAt) {
			// 已过期，删除节点
			this.deleteNode(node)
			delete(this.Map, key)
			this.len--
			return -1
		}

		// 未过期，刷新位置
		this.refresh(node)
		return node.val
	}
	return -1
}

// Put 插入或更新键值对，带 TTL
func (this *LRUCacheWithTTL) Put(key int, value int, ttl time.Duration) {
	this.mu.Lock()
	defer this.mu.Unlock()

	expireAt := time.Now().Add(ttl)

	// 检查 key 是否已存在
	if node, ok := this.Map[key]; ok {
		// key 已存在，更新值、过期时间并刷新位置
		node.val = value
		node.expireAt = expireAt
		this.refresh(node)
		return
	}

	// key 不存在，创建新节点
	node := &NodeWithTTL{
		key:      key,
		val:      value,
		expireAt: expireAt,
	}
	this.Map[key] = node
	this.len++
	this.refresh(node)

	if this.len > this.cap {
		// 删除尾节点（最久未使用）
		this.deleteTailNode()
		this.len--
	}
}

// cleanExpiredPeriodically 定期批量清理过期数据
func (this *LRUCacheWithTTL) cleanExpiredPeriodically() {
	for {
		select {
		case <-this.ticker.C:
			this.batchCleanExpired()
		case <-this.stopChan:
			this.ticker.Stop()
			return
		}
	}
}

// batchCleanExpired 批量清理过期节点
func (this *LRUCacheWithTTL) batchCleanExpired() {
	this.mu.Lock()
	defer this.mu.Unlock()

	now := time.Now()
	expiredKeys := make([]int, 0)

	// 遍历所有节点，收集过期的 key
	for key, node := range this.Map {
		if now.After(node.expireAt) {
			expiredKeys = append(expiredKeys, key)
		}
	}

	// 批量删除过期节点
	for _, key := range expiredKeys {
		if node, ok := this.Map[key]; ok {
			this.deleteNode(node)
			delete(this.Map, key)
			this.len--
		}
	}

	// 可选：记录清理日志
	if len(expiredKeys) > 0 {
		// fmt.Printf("Cleaned %d expired entries\n", len(expiredKeys))
	}
}

// Stop 停止后台清理 goroutine
func (this *LRUCacheWithTTL) Stop() {
	close(this.stopChan)
}

// deleteTailNode 删除尾节点
func (this *LRUCacheWithTTL) deleteTailNode() {
	node := this.Tail.Prev
	// 在 map 中删除
	delete(this.Map, node.key)

	// 在链表中删除
	this.deleteNode(node)
}

// deleteNode 从链表中删除节点
func (this *LRUCacheWithTTL) deleteNode(node *NodeWithTTL) {
	if node.Prev != nil && node.Next != nil {
		pre, nxt := node.Prev, node.Next
		pre.Next = nxt
		nxt.Prev = pre
	}
}

// refresh 将节点移到链表头部
func (this *LRUCacheWithTTL) refresh(node *NodeWithTTL) {
	// 如果节点已在链表中，先从原来的位置摘下来
	if node.Prev != nil && node.Next != nil {
		pre := node.Prev
		nxt := node.Next
		pre.Next = nxt
		nxt.Prev = pre
	}

	// 放到头部去
	node.Next = this.Head.Next
	node.Prev = this.Head

	this.Head.Next.Prev = node
	this.Head.Next = node
}

// Len 返回当前缓存中的元素数量
func (this *LRUCacheWithTTL) Len() int {
	this.mu.RLock()
	defer this.mu.RUnlock()
	return this.len
}
