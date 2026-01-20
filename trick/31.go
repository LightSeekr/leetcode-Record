package trick

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	//第一步：找“小数” (从后往前找第一个升序对)
	//从后往前遍历，找到第一个满足 nums[i] < nums[i+1] 的位置 i。
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 找到了 i
	if i >= 0 {
		//第二步：找“大数” (从后往前找第一个比 nums[i] 大的数)
		//在 i 后面的序列中，从后往前找，找到第一个比 nums[i] 大的数 nums[j]。
		j := n - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1, n-1)
}
func reverse(nums []int, start, end int) {
	for i, j := start, end; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i, j = i+1, j-1
	}
}
