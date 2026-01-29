package strings

/**
对于字符串 s 和 t，只有在 s = t + t + t + ... + t + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。

给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 x 能除尽 str2 。
*/

/*
*

 1. 如果存在公因子字符串，则 str1 + str2 == str2 + str1

 2. 公因子字符串的长度必定是两个字符串长度的最大公约数

 3. 使用辗转相除法求最大公约数

    时间复杂度：O(n + m)，空间复杂度：O(n + m)

*
*/
func gcdOfStrings(str1 string, str2 string) string {
	// 如果不存在公因子，直接返回空字符串
	if str1+str2 != str2+str1 {
		return ""
	}

	// 求两个字符串长度的最大公约数
	gcdLen := gcd(len(str1), len(str2))
	return str1[:gcdLen]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
