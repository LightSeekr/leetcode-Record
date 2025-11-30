package TwoPointers

func longestPalindrome(s string) string {
	// 针对每个子串都进行一个判断
	/**
	i 0 - len-1
	j i+1 - len-1
	判断 s[i,j] 是否是回文串
	这明显就是 O(n**3) 的想法
	*/
	length := 0
	res := ""
	for i := 0; i < len(s); i++ {
		r1 := palindrome(s, i, i)
		r2 := palindrome(s, i, i+1)

		if len(r1) >= len(r2) {
			if len(r1) > length {
				res = r1
				length = len(r1)
			}
		} else {
			if len(r2) > length {
				length = len(r2)
				res = r2
			}
		}
	}
	return res
}

// 找到中心点为某个位置的最长回文串
func palindrome(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}
