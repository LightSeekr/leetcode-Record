package binarySearch

func findMin(nums []int) int {
	n := len(nums)
	// n>=1
	// 1. 先找旋转的位置
	t1 := nums[n-1]
	// 旋转后
	//左半边大于 t1
	// 右半边 <=t1
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
	//全都小于 t1,那么 r =0
	//fmt.Printf("l: %d, r: %d\n", l, r)
	//fmt.Printf("nums[0,l]:%v\n", nums[0:l+1])
	//fmt.Printf("nums[r,n-1]:%v\n", nums[r:n])
	if r == 0 {
		// 说明数组未旋转
		// 升序数组
		return nums[0]
	} else {
		return nums[r]
	}
}

func FindMin(nums []int) int {
	return findMin(nums)
}
