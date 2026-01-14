package heap

/*
*
1. 用 map 记录每个数出现了多少次
2. 分桶。buckets[i] 存放所有出现次数为 i 的数字
3. 倒序收集：从 buckets 的最后（最高频率）往前遍历，把里面的数字拿出来，直到凑够 k 个。
*/
func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	buckets := make([][]int, n+1)
	for num, cnt := range countMap {
		buckets[cnt] = append(buckets[cnt], num)
	}

	res := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- {
		if len(buckets[i]) > 0 {
			res = append(res, buckets[i]...)
			if len(res) == k {
				break
			}
		}
	}
	return res
}

/**
我们可以维护一个大小为 k 的小顶堆。
- 堆里存的是结构体 (数值, 频率)。
- 我们根据频率来排序。
- 遍历所有元素，如果堆的大小超过 k，就把频率最小的（堆顶）弹出来，保留频率大的。
*/

// 定义堆中的节点
type Node struct {
	val   int // 数字本身
	count int // 出现的频率
}

func topKFrequent2(nums []int, k int) []int {
	// 1. 统计频率 - O(N)
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 2. 初始化一个小顶堆，容量为 k
	// 这里 len=0, cap=k，还是为了性能优化
	heap := make([]Node, 0, k+1)

	// 3. 遍历频率表，维护堆 - O(N log k)
	for val, count := range freqMap {
		node := Node{val, count}

		// 情况 A：堆还没满，直接加进去
		if len(heap) < k {
			heap = append(heap, node)
			siftUp(heap, len(heap)-1)
		} else {
			// 情况 B：堆满了，看看能不能踢掉堆顶（频率最小的）
			if count > heap[0].count {
				// 替换堆顶
				heap[0] = node
				// 重新调整堆结构
				siftDown(heap, 0, k)
			}
		}
	}

	// 4. 收集结果
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap[i].val
	}
	return res
}

// --- 手写堆的核心函数 ---

// 上浮操作：当向末尾添加新元素时调用
func siftUp(heap []Node, i int) {
	for {
		parent := (i - 1) / 2
		// 如果到了根节点，或者父节点比自己小（符合小顶堆规则），停止
		if i == 0 || heap[parent].count <= heap[i].count {
			break
		}
		// 否则交换，继续往上找
		heap[parent], heap[i] = heap[i], heap[parent]
		i = parent
	}
}

// 下沉操作：当替换堆顶时调用
func siftDown(heap []Node, i, n int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i // 假设当前节点最小

		// 找左孩子比一比
		if left < n && heap[left].count < heap[smallest].count {
			smallest = left
		}
		// 找右孩子比一比
		if right < n && heap[right].count < heap[smallest].count {
			smallest = right
		}

		// 如果最小的还是自己，说明位置对了，不用沉了
		if smallest == i {
			break
		}

		// 交换，继续往下沉
		heap[i], heap[smallest] = heap[smallest], heap[i]
		i = smallest
	}
}
