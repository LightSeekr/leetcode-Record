package hash

/*
*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
还是要用 O(n) 空间的
使用hash 标记才可以
*/
func firstMissingPositive(nums []int) int {
	Map := make(map[int]int)
	for _, num := range nums {
		if num > 0 {
			Map[num]++
		}

	}
	res := 1
	for {
		if _, ok := Map[res]; ok {
			res++
		} else {
			return res
		}
	}
}
