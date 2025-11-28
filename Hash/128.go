package Hash

const (
	TenToThe9th = 1000000000
	Length      = TenToThe9th*2 + 1
)

// 这种方法可以做，但是在 leetcode 上会 panic 因为数组开太大了。
func longestConsecutive2(nums []int) int {
	// 边界 case 特判
	if len(nums) == 0 {
		return 0
	}
	// 用数组做哈希表该数组的长度为 10^9*2+1 个数
	exist := [Length]int64{}
	// 遍历第一遍 存在为 1 不存在为 0 特殊处理，每一个数都加上 10^9 保证从0 开始
	for _, num := range nums {
		exist[num+TenToThe9th]++
	}
	// 遍历第二遍，找出连续个 1的最长长度
	res := 0
	st, ed := 0, 0
	for i := 0; i < Length; {
		if exist[i] >= 1 {
			st = i
			ed = i + 1
			for exist[ed] >= 1 {
				ed++
			}
			cnt := ed - st
			if cnt > res {
				res = cnt
			}
			i = ed + 1
		} else {
			i++
		}
	}
	return res
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	res := 0
	for num := range numSet {
		if _, ok := numSet[num-1]; ok {
			continue
		}
		cur := num
		curlen := 1
		for numSet[cur+1] {
			cur++
			curlen++
		}
		if curlen > res {
			res = curlen
		}
	}
	return res

}
func LongestConsecutive(nums []int) int {
	return longestConsecutive(nums)
}
