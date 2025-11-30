package TwoPointers

func moveZeroes(nums []int) {
	// 要求保持非零元素的相对顺序。
	i, j := 0, 0 // i 为合理数组的末尾,j 探路
	for j < len(nums) {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	//把后面的元素 =0
	for i < len(nums) {
		nums[i] = 0
		i++
	}
}

func MoveZeroes(nums []int) {
	moveZeroes(nums)
}
