package enumerate

/*
*
给你一个 不包含 任何零的整数数组 nums ，找出自身与对应的负数都在数组中存在的最大正整数 k 。
返回正整数 k ，如果不存在这样的整数，返回 -1 。
*/

/*
*
放入 map 的一定是原来本身的数字
*/
func findMaxK(nums []int) int {
	cntMap := make(map[int]bool)
	res := -1
	for _, n := range nums {
		// 对应的数存在
		if _, ok := cntMap[n*-1]; ok {
			if n < 0 {
				res = max(res, n*-1)
			} else {
				res = max(res, n)
			}
		}
		cntMap[n] = true
	}
	return res
}

func FindMaxK(nums []int) int {
	return findMaxK(nums)
}
