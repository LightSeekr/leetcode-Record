package slideWindow

/*
*
给定两个字符串 s 和 t，长度分别是 m 和 n，返回 s 中的 最短窗口 子串，使得该子串包含 t 中的每一个字符（包括重复字符）。如果没有这样的子串，返回空字符串 ""。

测试用例保证答案唯一。

已按 76.go 注释的题意把 minWindow(s, t) 补全了：用滑动窗口在 s 里找最短子串，覆盖 t 的所有字符（包含重复次数），找

	不到就返回 ""。

	- 实现位置：76.go:8
	- 核心思路：need[256] 统计还需要的字符数，remain 统计还差多少个字符没覆盖；右指针扩张窗口，remain==0 时左指针尽量收
	  缩并更新最短答案。
	- 说明：当前实现按 byte 处理（适合 LeetCode 76 的常见输入场景）；如果你要支持完整 Unicode，需要改成按 rune +
	  map[rune]int 计数。
*/
func minWindow(s string, t string) string {
	// LeetCode 76: Minimum Window Substring
	// 用滑动窗口在 s 中找最短子串，使其包含 t 中所有字符（含重复次数）。
	if len(t) == 0 || len(s) == 0 || len(s) < len(t) {
		return ""
	}

	need := [256]int{}
	remain := len(t) // 还需要匹配的字符总数（含重复）
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	bestL, bestR := 0, 0
	bestLen := len(s) + 1

	l := 0
	for r := 0; r < len(s); r++ {
		c := s[r]
		if need[c] > 0 {
			remain--
		}
		need[c]--

		// 当窗口已经覆盖 t，尽量收缩左边界
		for remain == 0 {
			if r-l+1 < bestLen {
				bestLen = r - l + 1
				bestL, bestR = l, r+1 // r+1 方便切片
			}

			cl := s[l]
			need[cl]++
			if need[cl] > 0 {
				remain++
			}
			l++
		}
	}

	if bestLen == len(s)+1 {
		return ""
	}
	return s[bestL:bestR]
}
