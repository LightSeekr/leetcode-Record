package dynamicProgramming

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	maxLen := 1
	begin := 0

	// 注意：DP 的遍历顺序很重要
	// 必须先计算短的子串，再计算长的
	// 所以外层循环枚举长度 L，或者倒序枚举 i

	// 写法：倒序枚举左边界 i
	for i := n - 1; i >= 0; i-- {
		// 枚举右边界 j
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				// 如果长度 <= 2 (即 j-i < 2，如 "aa", "aba")，直接 true
				// 否则看内部 dp[i+1][j-1]
				if j-i <= 2 || dp[i+1][j-1] {
					dp[i][j] = true

					// 记录最长
					if j-i+1 > maxLen {
						maxLen = j - i + 1
						begin = i
					}
				}
			}
		}
	}

	return s[begin : begin+maxLen]
}
