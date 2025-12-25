package hash

// 哈希做法
func twoSum(nums []int, target int) []int {
	res := []int{0, 0}
	ma := make(map[int]int) //val-> idx 的映射

	for idx, val := range nums {
		if idx1, ok := ma[target-val]; ok {
			res[0] = idx
			res[1] = idx1
			return res
		}
		ma[val] = idx
	}
	return res
}

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}
