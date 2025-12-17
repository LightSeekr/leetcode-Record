package BinarySearch

import "fmt"

/*
*二分不一定要有序，而是满足单调
 */
func search(nums []int, target int) int {
	n := len(nums)
	// 1. 先找旋转的位置
	// 2. 再找 target 是否存在
	t1 := nums[n-1]
	// 旋转后
	//左半边大于 target
	// 右半边 <=target
	l, r := -1, n
	for r-l > 1 {
		mid := l + (r-l)/2
		if nums[mid] <= t1 {
			r = mid
		} else {
			l = mid
		}
	}
	// l,r 是边界
	// l 是>t1 的右边界
	// r 是<=t1 的左边界
	// 如果没有旋转，会发生什么？
	//全都小于 t1,那么 r =0
	// [0,l] [r,n-1] 是两个数组了
	fmt.Printf("l: %d, r: %d\n", l, r)
	fmt.Printf("nums[0,l]:%v\n", nums[0:l+1])
	fmt.Printf("nums[r,n-1]:%v\n", nums[r:n])
	if r == 0 {
		// 说明数组未旋转，在 [0,n) 中找 target
		r1 := binarySearch(nums[0:n], target)
		return r1
	} else {
		if target > nums[n-1] {
			//在 [0,l] 中找 target
			r1 := binarySearch(nums[0:l+1], target)
			return r1
		} else if target < nums[n-1] {
			r1 := binarySearch(nums[r:n], target)
			if r1 != -1 {
				// 这里需要转一下，求的是原来的数组的下标
				return r + r1
			} else {
				return r1
			}
		} else {
			return n - 1
		}

	}
}

func binarySearch(nums []int, target int) int {
	n := len(nums)
	l, r := -1, n
	for r-l > 1 {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid
		}
	}
	// l,r 是边界
	// l 是<target 的右边界
	// r 是 >=target 的左边界
	//return r
	if r == n {
		// 说明数组中不存在这个数
		return -1
	} else {
		if nums[r] == target {
			return r
		} else {
			return -1
		}
	}
}

func Search(nums []int, target int) int {
	return search(nums, target)
}
