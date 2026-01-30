package LRUCache

import (
	"testing"
	"time"
)

func TestLRUCacheWithTTL_BasicOperations(t *testing.T) {
	// 创建容量为 2 的缓存，每 1 秒清理一次
	cache := ConstructorWithTTL(2, 1*time.Second)
	defer cache.Stop()

	// 测试基本的 Put 和 Get
	cache.Put(1, 100, 5*time.Second)
	cache.Put(2, 200, 5*time.Second)

	if val := cache.Get(1); val != 100 {
		t.Errorf("Expected 100, got %d", val)
	}

	if val := cache.Get(2); val != 200 {
		t.Errorf("Expected 200, got %d", val)
	}
}

func TestLRUCacheWithTTL_PassiveExpiration(t *testing.T) {
	cache := ConstructorWithTTL(2, 10*time.Second)
	defer cache.Stop()

	// 插入一个 1 秒后过期的数据
	cache.Put(1, 100, 1*time.Second)

	// 立即获取，应该成功
	if val := cache.Get(1); val != 100 {
		t.Errorf("Expected 100, got %d", val)
	}

	// 等待 1.5 秒后再获取，应该返回 -1（被动过期）
	time.Sleep(1500 * time.Millisecond)
	if val := cache.Get(1); val != -1 {
		t.Errorf("Expected -1 (expired), got %d", val)
	}

	// 验证缓存长度已减少
	if cache.Len() != 0 {
		t.Errorf("Expected cache length 0, got %d", cache.Len())
	}
}

func TestLRUCacheWithTTL_ActiveCleaning(t *testing.T) {
	// 创建缓存，每 500ms 清理一次
	cache := ConstructorWithTTL(3, 500*time.Millisecond)
	defer cache.Stop()

	// 插入 3 个数据，都是 1 秒后过期
	cache.Put(1, 100, 1*time.Second)
	cache.Put(2, 200, 1*time.Second)
	cache.Put(3, 300, 1*time.Second)

	if cache.Len() != 3 {
		t.Errorf("Expected cache length 3, got %d", cache.Len())
	}

	// 等待 1.5 秒，让定期清理触发
	time.Sleep(1500 * time.Millisecond)

	// 此时后台清理应该已经删除了所有过期数据
	if cache.Len() != 0 {
		t.Errorf("Expected cache length 0 after cleaning, got %d", cache.Len())
	}
}

func TestLRUCacheWithTTL_LRUEviction(t *testing.T) {
	cache := ConstructorWithTTL(2, 10*time.Second)
	defer cache.Stop()

	// 插入 2 个数据
	cache.Put(1, 100, 10*time.Second)
	cache.Put(2, 200, 10*time.Second)

	// 插入第 3 个数据，应该淘汰最久未使用的 key=1
	cache.Put(3, 300, 10*time.Second)

	if val := cache.Get(1); val != -1 {
		t.Errorf("Expected -1 (evicted), got %d", val)
	}

	if val := cache.Get(2); val != 200 {
		t.Errorf("Expected 200, got %d", val)
	}

	if val := cache.Get(3); val != 300 {
		t.Errorf("Expected 300, got %d", val)
	}
}

func TestLRUCacheWithTTL_UpdateExisting(t *testing.T) {
	cache := ConstructorWithTTL(2, 10*time.Second)
	defer cache.Stop()

	// 插入数据，1 秒后过期
	cache.Put(1, 100, 1*time.Second)

	// 等待 0.5 秒
	time.Sleep(500 * time.Millisecond)

	// 更新数据，延长过期时间到 2 秒
	cache.Put(1, 200, 2*time.Second)

	// 再等待 1 秒（总共 1.5 秒），原本应该过期，但因为更新了 TTL，所以还有效
	time.Sleep(1 * time.Second)

	if val := cache.Get(1); val != 200 {
		t.Errorf("Expected 200, got %d", val)
	}

	// 再等待 1.5 秒（总共 3 秒），现在应该过期了
	time.Sleep(1500 * time.Millisecond)

	if val := cache.Get(1); val != -1 {
		t.Errorf("Expected -1 (expired), got %d", val)
	}
}

func TestLRUCacheWithTTL_MixedTTL(t *testing.T) {
	cache := ConstructorWithTTL(3, 500*time.Millisecond)
	defer cache.Stop()

	// 插入不同 TTL 的数据
	cache.Put(1, 100, 1*time.Second) // 1 秒后过期
	cache.Put(2, 200, 2*time.Second) // 2 秒后过期
	cache.Put(3, 300, 3*time.Second) // 3 秒后过期

	// 等待 1.5 秒
	time.Sleep(1500 * time.Millisecond)

	// key=1 应该已经被清理
	if val := cache.Get(1); val != -1 {
		t.Errorf("Expected -1 (expired), got %d", val)
	}

	// key=2 和 key=3 应该还有效
	if val := cache.Get(2); val != 200 {
		t.Errorf("Expected 200, got %d", val)
	}

	if val := cache.Get(3); val != 300 {
		t.Errorf("Expected 300, got %d", val)
	}

	// 等待 1 秒（总共 2.5 秒）
	time.Sleep(1 * time.Second)

	// key=2 应该过期
	if val := cache.Get(2); val != -1 {
		t.Errorf("Expected -1 (expired), got %d", val)
	}

	// key=3 应该还有效
	if val := cache.Get(3); val != 300 {
		t.Errorf("Expected 300, got %d", val)
	}
}

// 性能测试：验证定期清理不会阻塞正常操作
func TestLRUCacheWithTTL_ConcurrentAccess(t *testing.T) {
	cache := ConstructorWithTTL(100, 100*time.Millisecond)
	defer cache.Stop()

	// 并发写入
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 100; j++ {
				cache.Put(id*100+j, j, 2*time.Second)
			}
			done <- true
		}(i)
	}

	// 等待所有写入完成
	for i := 0; i < 10; i++ {
		<-done
	}

	// 等待一段时间让清理运行
	time.Sleep(500 * time.Millisecond)

	// 验证缓存仍然可以正常工作
	cache.Put(9999, 9999, 5*time.Second)
	if val := cache.Get(9999); val != 9999 {
		t.Errorf("Expected 9999, got %d", val)
	}
}
