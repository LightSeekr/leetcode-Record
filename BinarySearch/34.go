package BinarySearch

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	l, r := -1, n
	for r-l > 1 {
		mid := l + (r-l)/2

		if nums[mid] >= target {
			r = mid
		} else {
			l = mid
		}

	}
	// l,r是边界
	// 存在全都小于 target 的情况 所以 r>=n
	if r >= n || nums[r] != target {
		// 不存在 target
		return []int{-1, -1}
	} else {
		// r是>=target 的左边界
		// l 是<target 的有边界
		idx := r + 1
		for idx < n && nums[idx] == target {
			idx++
		}
		return []int{r, idx - 1}
	}
}

func SearchRange(nums []int, target int) []int {
	return searchRange(nums, target)
}
