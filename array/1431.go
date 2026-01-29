package array

/**
有 n 个有糖果的孩子。给你一个数组 candies，其中 candies[i] 代表第 i 个孩子拥有的糖果数目，和一个整数 extraCandies 表示你所有的额外糖果的数量。

返回一个长度为 n 的布尔数组 result，如果把所有的 extraCandies 给第 i 个孩子之后，他会拥有所有孩子中 最多 的糖果，那么 result[i] 为 true，否则为 false。
注意，允许有多个孩子同时拥有 最多 的糖果数目。
*/

/*
*

 1. 先遍历一次找出当前糖果数的最大值

 2. 再遍历一次，判断每个孩子加上额外糖果后是否 >= 最大
 2. 再遍历一次，判断每个孩子加上额外糖果后是否 >= 最大值
    时间复杂度：O(n)，空间复杂度：O(n)（结果数组）

*
*/
**/
func kidsWithCandies(candies []int, extraCandies int) []bool {
	// 找出当前最大糖果数
	maxCandies := 0
	for _, c := range candies {
		if c > maxCandies {
			maxCandies = c
		}
	}

	// 判断每个孩子加上额外糖果后是否能成为最多
	result := make([]bool, len(candies))
	for i, c := range candies {
		result[i] = c+extraCandies >= maxCandies
	}

	return result
}
