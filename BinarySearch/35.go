package BinarySearch

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

找 >= target 的左边界
如果 num 全都大于 target 那么 r 应该放入 target
*/

func searchInsert(nums []int, target int) int {
	l, r, mid := -1, len(nums), 0
	for r-l > 1 {
		mid = l + (r-l)/2
		//fmt.Printf("l: %d, r: %d, mid: %d\n", l, r, mid)
		if nums[mid] >= target {
			r = mid
		} else if nums[mid] < target {
			l = mid
		}
	}
	return r
}

func SearchInsert(nums []int, target int) int {
	return searchInsert(nums, target)
}
