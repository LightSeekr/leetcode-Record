package dynamicProgramming

func longestCommonSubsequence(text1 string, text2 string) int {
	row, col := len(text1), len(text2)
	dp := make([][]int, row+1)
	for i := range dp {
		dp[i] = make([]int, col+1)
	}
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[row][col]
}
