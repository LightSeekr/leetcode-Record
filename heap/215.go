package heap

import "math/rand"

func findKthLargest(nums []int, k int) int {
	// 使用快速选择算法
	// 第 k 大元素的索引是 k-1（从大到小排序）
	return quickSelect(nums, 0, len(nums)-1, k-1)
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	// 随机选择 pivot 并分区
	pivotIndex := partition(nums, left, right)

	if pivotIndex == k {
		return nums[pivotIndex]
	} else if pivotIndex < k {
		// 第 k 大在右边
		return quickSelect(nums, pivotIndex+1, right, k)
	} else {
		// 第 k 大在左边
		return quickSelect(nums, left, pivotIndex-1, k)
	}
}

func partition(nums []int, left, right int) int {
	randomIndex := left + rand.Intn(right-left+1)
	nums[randomIndex], nums[right] = nums[right], nums[randomIndex]

	// 选择最右边的元素作为 pivot
	pivot := nums[right]
	i := left

	// 将大于 pivot 的元素放到左边
	for j := left; j < right; j++ {
		if nums[j] > pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	// 将 pivot 放到正确位置
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// findKthLargestWithHeap 使用小顶堆解法
func findKthLargestWithHeap(nums []int, k int) int {
	// 使用最小堆，维护 k 个最大的元素
	// 堆顶就是第 k 大的元素
	h := &MinHeap{}

	for _, num := range nums {
		if h.Len() < k {
			heapPush(h, num)
		} else if num > (*h)[0] {
			heapPop(h)
			heapPush(h, num)
		}
	}

	return (*h)[0]
}

// MinHeap 最小堆实现
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// heapPush 向堆中添加元素
func heapPush(h *MinHeap, x int) {
	h.Push(x)
	up(h, h.Len()-1)
}

// heapPop 从堆中弹出最小元素
func heapPop(h *MinHeap) int {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop().(int)
}

// up 上浮操作
func up(h *MinHeap, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

// down 下沉操作
func down(h *MinHeap, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
