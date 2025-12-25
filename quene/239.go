package quene

/*
*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。

https://www.bilibili.com/video/BV1bM411X72E/?vd_source=4ee9e5a84cfef330f8fb10409e1fdfb1
单调队列的形成过程。
为什么说其是单调的。

	maxSlidingWindow 用单调队列（更准确说：单调双端队列 / monotonic deque）的原因只有一个：在窗口每次右移一格的过程中，
	 把“每个窗口求最大值”从 O(k) 降到摊还 O(1)，整体做到 O(n)。

	 如果不用单调队列，最直接的做法是：

	 - 每个窗口都扫描 k 个元素取 max → 总复杂度 O(n*k)，k 大时会超时。

	 单调队列的想法是：维护一个“候选最大值列表”，保证队头永远是当前窗口最大值。

	 为什么它能做到这点

	 - 队列里存的是下标（不是值），这样才能判断元素是否“滑出窗口”（下标 < 左边界就得踢掉）。
	 - 队列从头到尾对应的 nums[下标] 是“单调递减”的：
	     - 新元素 nums[i] 进来时，把队尾所有 <= nums[i] 的下标都弹出
	         - 因为它们既更小、又更早过期，未来任何窗口里都不可能赢过 nums[i]，留着只会干扰。
	     - 然后把 i 加到队尾。
	 - 窗口左边界变成 left = i-k+1 时：
	     - 如果队头下标 < left，说明最大值已经不在窗口里了，弹出队头。
	 - 此时队头就是当前窗口最大值的下标，答案直接取 nums[deq[0]]。

	 为什么说是“摊还 O(1)”
	 每个下标：

	 - 进队一次（push）
	 - 出队最多一次（要么从队头滑出窗口，要么在某次入队时从队尾被更大的元素“碾掉”）
	   所以总弹出次数是 O(n)，整体就是 O(n)。
*/
func maxSlidingWindow(nums []int, k int) []int {
	if k <= 0 || len(nums) == 0 || k > len(nums) {
		return []int{}
	}
	// 单调队列（存下标），队头永远是当前窗口最大值的下标。
	deque := make([]int, 0, len(nums))
	res := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		// 1) 把窗口左边界之外的下标弹出（i-k 是刚刚离开窗口的位置）。
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// 2) 维持队列单调递减：新元素 nums[i] 比队尾大（或相等）时，队尾不可能再成为最大值。
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)

		// 3) 当窗口形成（i >= k-1）后，记录当前最大值。
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}
