package hash

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	// 这题先用哈希来做
	Map := make(map[int]int)
	for idx1, val := range nums {
		if idx2, ok := Map[val]; ok {
			if math.Abs(float64(idx2-idx1)) <= float64(k) {
				return true
			}
		}
		Map[val] = idx1
	}
	return false
}

func ContainsNearbyDuplicate(nums []int, k int) bool {
	return containsNearbyDuplicate(nums, k)
}
