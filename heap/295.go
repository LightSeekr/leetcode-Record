package heap

type MedianFinder struct {
	small []int // 大顶堆：存较小的一半
	large []int // 小顶堆：存较大的一半
}

func Constructor() MedianFinder {
	return MedianFinder{
		small: []int{},
		large: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 1. 先放入 small (大顶堆)
	this.small = append(this.small, num)
	// 上浮调整，cmp 为 > 表示大顶堆
	siftup(this.small, func(a, b int) bool { return a > b })

	// 2. 把 small 的堆顶（这一半里最大的）移到 large
	maxVal := this.small[0]
	// 删除 small 堆顶
	this.small = popHeap(this.small, func(a, b int) bool { return a > b })

	// 放入 large (小顶堆)
	this.large = append(this.large, maxVal)
	// 上浮调整，cmp 为 < 表示小顶堆
	siftup(this.large, func(a, b int) bool { return a < b })

	// 3. 维持平衡：约定 len(small) >= len(large)
	// 如果 large 比 small 多了，把 large 的堆顶还给 small
	if len(this.large) > len(this.small) {
		minVal := this.large[0]
		this.large = popHeap(this.large, func(a, b int) bool { return a < b })

		this.small = append(this.small, minVal)
		siftup(this.small, func(a, b int) bool { return a > b })
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.small) > len(this.large) {
		return float64(this.small[0])
	}
	return float64(this.small[0]+this.large[0]) / 2.0
}

// --- 手写堆的核心辅助函数 ---

// 比较函数类型：func(a, b) bool
// 如果是小顶堆，传入 <；如果是大顶堆，传入 >
type Compare func(int, int) bool

// 上浮 (Push 后调用)
func siftup(heap []int, cmp Compare) {
	i := len(heap) - 1
	for {
		parent := (i - 1) / 2
		// 如果到了根，或者父节点满足堆序 (parent "better than" i)，停止
		if i == 0 || cmp(heap[parent], heap[i]) {
			break
		}
		heap[parent], heap[i] = heap[i], heap[parent]
		i = parent
	}
}

// 下沉 (Pop 后调用)
func siftdown(heap []int, i int, cmp Compare) {
	n := len(heap)
	for {
		left := 2*i + 1
		right := 2*i + 2
		best := i // 假设当前节点是“最符合条件”的（最大或最小）

		if left < n && cmp(heap[left], heap[best]) {
			best = left
		}
		if right < n && cmp(heap[right], heap[best]) {
			best = right
		}

		if best == i {
			break
		}
		heap[i], heap[best] = heap[best], heap[i]
		i = best
	}
}

// 辅助 Pop 函数：交换堆顶和末尾，删除末尾，然后下沉堆顶
// 返回新的切片
func popHeap(heap []int, cmp Compare) []int {
	n := len(heap)
	// 交换堆顶和最后一个元素
	heap[0], heap[n-1] = heap[n-1], heap[0]
	// 切掉最后一个
	heap = heap[:n-1]
	// 如果还没空，下沉新的堆顶
	if len(heap) > 0 {
		siftdown(heap, 0, cmp)
	}
	return heap
}
