package enumerate

/*
*
如果一组数字 (i,j) 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对 。
*/
func numIdenticalPairs(nums []int) int {
	cntMap := make(map[int]int)
	res := 0
	for _, n := range nums {
		if _, ok := cntMap[n]; ok {
			res += cntMap[n] // 说明前面有 cntMap[n] 个数与 n 相等
		}
		cntMap[n] = cntMap[n] + 1
	}
	return res
}
func NumIdenticalPairs(nums []int) int {
	return numIdenticalPairs(nums)
}
