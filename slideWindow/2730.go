package slideWindow

func longestSemiRepetitiveSubstring(s string) int {
	ns := len(s)
	if ns < 2 {
		return ns
	}
	//请你返回 s 中最长 半重复 子字符串 的长度。
	same, l, r, ans := 0, 0, 0, 0
	for r = 1; r < ns; r++ {
		if s[r] == s[r-1] {
			same++
		}
		if same > 1 {
			l++
			for s[l] != s[l-1] {
				l++
			}
			same = 1
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

func LongestSemiRepetitiveSubstring(s string) int {
	return longestSemiRepetitiveSubstring(s)
}
