package stack

func RemoveDuplicates(nums []int, n int) int {
	l := len(nums)
	if l == 0 || l == 1 {
		return l
	}
	return process(nums, n)
}

func process(nums []int, n int) int {
	l := len(nums)
	k := n - 1
	for i := n; i < l; i++ {
		if nums[i] != nums[k-n+1] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}
