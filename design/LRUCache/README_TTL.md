# 带超时时间的 LRU 缓存实现

## 核心设计：混合方案

结合了**被动过期**和**主动清理**两种策略，类似 Redis 的过期键删除策略。

## 工作原理

### 1. 被动过期（Lazy Expiration）

在 `Get` 操作时检查数据是否过期：

```go
func (this *LRUCacheWithTTL) Get(key int) int {
    if node, ok := this.Map[key]; ok {
        // 检查是否过期
        if time.Now().After(node.expireAt) {
            // 立即删除过期数据
            this.deleteNode(node)
            delete(this.Map, key)
            this.len--
            return -1
        }
        // 未过期，正常返回
        this.refresh(node)
        return node.val
    }
    return -1
}
```

**优点**：
- 保证返回数据的准确性
- 不会返回已过期的数据
- 实现简单

**缺点**：
- 如果某个 key 永远不被访问，过期数据会一直占用内存

---

### 2. 主动清理（Active Expiration）

使用后台 goroutine 定期批量清理过期数据：

#### 2.1 启动定期清理

```go
func ConstructorWithTTL(capacity int, cleanInterval time.Duration) *LRUCacheWithTTL {
    // ... 初始化代码

    l.ticker = time.NewTicker(cleanInterval)  // 创建定时器
    go l.cleanExpiredPeriodically()           // 启动后台清理

    return l
}
```

#### 2.2 定期清理循环

```go
func (this *LRUCacheWithTTL) cleanExpiredPeriodically() {
    for {
        select {
        case <-this.ticker.C:
            // 定时器触发，执行批量清理
            this.batchCleanExpired()
        case <-this.stopChan:
            // 收到停止信号，退出 goroutine
            this.ticker.Stop()
            return
        }
    }
}
```

#### 2.3 批量清理逻辑

```go
func (this *LRUCacheWithTTL) batchCleanExpired() {
    this.mu.Lock()
    defer this.mu.Unlock()

    now := time.Now()
    expiredKeys := make([]int, 0)

    // 第一步：遍历所有节点，收集过期的 key
    for key, node := range this.Map {
        if now.After(node.expireAt) {
            expiredKeys = append(expiredKeys, key)
        }
    }

    // 第二步：批量删除过期节点
    for _, key := range expiredKeys {
        if node, ok := this.Map[key]; ok {
            this.deleteNode(node)
            delete(this.Map, key)
            this.len--
        }
    }
}
```

**为什么分两步？**
- 避免在遍历 map 时修改 map（Go 中这样做会 panic）
- 先收集，再删除，逻辑更清晰

---

## 关键设计点

### 1. 并发安全

使用 `sync.RWMutex` 保护共享数据：

```go
type LRUCacheWithTTL struct {
    // ...
    mu sync.RWMutex  // 读写锁
}

func (this *LRUCacheWithTTL) Get(key int) int {
    this.mu.Lock()
    defer this.mu.Unlock()
    // ...
}
```

- **写操作**（Put, 清理）：使用 `Lock()`
- **读操作**（Len）：可以使用 `RLock()` 提高并发性能

### 2. 清理间隔的选择

```go
// 示例：不同场景的清理间隔
cache1 := ConstructorWithTTL(100, 1*time.Second)   // 高频清理
cache2 := ConstructorWithTTL(100, 10*time.Second)  // 低频清理
cache3 := ConstructorWithTTL(100, 100*time.Millisecond) // 极高频清理
```

**选择建议**：
- **高频访问场景**：1-5 秒，依赖被动过期为主
- **低频访问场景**：100-500 毫秒，主动清理更重要
- **内存敏感场景**：更短的间隔，及时释放内存

### 3. 优雅关闭

```go
func (this *LRUCacheWithTTL) Stop() {
    close(this.stopChan)  // 通知后台 goroutine 退出
}

// 使用时
cache := ConstructorWithTTL(100, 1*time.Second)
defer cache.Stop()  // 确保 goroutine 被清理
```

---

## 使用示例

### 基本使用

```go
package main

import (
    "fmt"
    "time"
    "your-project/leetcode/design/LRUCache"
)

func main() {
    // 创建容量为 3 的缓存，每 1 秒清理一次
    cache := LRUCache.ConstructorWithTTL(3, 1*time.Second)
    defer cache.Stop()

    // 插入数据，5 秒后过期
    cache.Put(1, 100, 5*time.Second)
    cache.Put(2, 200, 5*time.Second)

    // 立即获取
    fmt.Println(cache.Get(1))  // 输出: 100

    // 等待 6 秒后获取
    time.Sleep(6 * time.Second)
    fmt.Println(cache.Get(1))  // 输出: -1 (已过期)
}
```

### 不同 TTL 的数据

```go
cache := LRUCache.ConstructorWithTTL(10, 500*time.Millisecond)
defer cache.Stop()

// 短期缓存：1 秒
cache.Put(1, 100, 1*time.Second)

// 中期缓存：10 秒
cache.Put(2, 200, 10*time.Second)

// 长期缓存：1 分钟
cache.Put(3, 300, 60*time.Second)
```

### 更新 TTL

```go
// 第一次插入，5 秒后过期
cache.Put(1, 100, 5*time.Second)

// 更新值和 TTL，重新计时 10 秒
cache.Put(1, 200, 10*time.Second)
```

---

## 性能分析

### 时间复杂度

| 操作 | 时间复杂度 | 说明 |
|------|-----------|------|
| Get | O(1) | HashMap + 双向链表 |
| Put | O(1) | HashMap + 双向链表 |
| 批量清理 | O(n) | 遍历所有节点，n 为缓存大小 |

### 空间复杂度

- **额外空间**：O(n)，n 为缓存容量
- **每个节点**：增加一个 `time.Time` 字段（24 字节）

### 清理开销

假设缓存容量为 1000，清理间隔为 1 秒：

- **最坏情况**：每秒遍历 1000 个节点
- **实际开销**：通常很小，因为大部分节点未过期
- **优化方案**：可以采样清理（每次只检查部分节点）

---

## 优化方向

### 1. 采样清理（类似 Redis）

```go
func (this *LRUCacheWithTTL) batchCleanExpired() {
    this.mu.Lock()
    defer this.mu.Unlock()

    now := time.Now()
    sampleSize := 20  // 每次只检查 20 个节点
    checked := 0

    for key, node := range this.Map {
        if checked >= sampleSize {
            break
        }

        if now.After(node.expireAt) {
            this.deleteNode(node)
            delete(this.Map, key)
            this.len--
        }
        checked++
    }
}
```

### 2. 使用最小堆

维护一个按过期时间排序的最小堆，只清理堆顶的过期节点：

```go
// 伪代码
func (this *LRUCacheWithTTL) batchCleanExpired() {
    now := time.Now()

    // 只清理堆顶的过期节点
    for this.expireHeap.Len() > 0 {
        node := this.expireHeap.Peek()
        if now.Before(node.expireAt) {
            break  // 堆顶未过期，后面的也不会过期
        }

        this.expireHeap.Pop()
        this.deleteNode(node)
        delete(this.Map, node.key)
        this.len--
    }
}
```

**优点**：清理时间复杂度降低到 O(k log n)，k 为过期节点数

### 3. 分层时间轮

将不同过期时间的数据放入不同的时间槽，类似 Kafka 的时间轮算法。

---

## 与 Redis 的对比

| 特性 | 本实现 | Redis |
|------|--------|-------|
| 被动过期 | ✅ Get 时检查 | ✅ 访问时检查 |
| 定期清理 | ✅ 全量遍历 | ✅ 采样清理（每次 20 个） |
| 最大内存淘汰 | ✅ LRU | ✅ 多种策略（LRU/LFU/Random） |
| 并发安全 | ✅ RWMutex | ✅ 单线程 + IO 多路复用 |

---

## 测试

运行测试：

```bash
cd /Users/yolo/Program/go_interview/leetcode/design/LRUCache
go test -v -run TestLRUCacheWithTTL
```

测试覆盖：
- ✅ 基本操作（Put/Get）
- ✅ 被动过期
- ✅ 主动清理
- ✅ LRU 淘汰
- ✅ 更新 TTL
- ✅ 混合 TTL
- ✅ 并发访问

---

## 总结

这个实现展示了如何在 LRU 缓存的基础上添加 TTL 功能，核心要点：

1. **混合策略**：被动过期 + 主动清理，平衡性能和内存
2. **定期清理**：使用 `time.Ticker` 和 goroutine 实现后台清理
3. **并发安全**：使用 `sync.RWMutex` 保护共享数据
4. **优雅关闭**：通过 channel 通知后台 goroutine 退出

这种设计在面试中可以很好地展示你对：
- 数据结构的理解（双向链表 + HashMap）
- 并发编程的掌握（goroutine + channel + mutex）
- 系统设计的思考（过期策略、性能优化）
