package simulation

import (
	"fmt"
	"sort"
)

/*
*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
要排序吗？
对一维数组排序，
starti 小的排在前面
如果 starti 小，那么 endi 小的排在前面
排序，然后比对
和双指针没关系，排序就好了。
*/
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		// 首先比较 start (索引 0)
		if intervals[i][0] != intervals[j][0] {
			// start 小的排在前面 (升序)
			return intervals[i][0] < intervals[j][0]
		}
		// 如果 start 相同，则比较 end (索引 1)
		// end 小的排在前面 (升序)
		return intervals[i][1] < intervals[j][1]
	})
	fmt.Println("intervals: ", intervals)
	res := make([][]int, 0)
	// 控制进展
	for i := 0; i < len(intervals); {
		l, r := intervals[i][0], intervals[i][1] // 每一轮合并的结果
		j := 0
		for j = i + 1; j < len(intervals); j++ {
			// 在排序后，j>i
			// 说明j 和 i 两个区间无交集
			if intervals[j][0] > r {
				break
			}
			// 说明有交集
			// intervals[j][0] ==intervals[i][0]
			// intervals[j][1] >=intervals[i][1]
			// 合并，更新区间
			l = min(l, intervals[j][0])
			r = max(intervals[j][1], r)
		}
		res = append(res, []int{l, r})
		// j 就是下一个 i 的位置
		i = j
	}
	return res
}
