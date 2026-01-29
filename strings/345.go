package strings

/**
给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。

元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。
*/
/**
 1. 使用双指针，左指针从头开始，右指针从尾开始
 2. 左指针找到元音字母，右指针找到元音字母，然后交换
 3. 重复直到左右指针相遇

 时间复杂度：O(n)，空间复杂度：O(n)
**/
func reverseVowels(s string) string {
	bytes := []byte(s)
	left, right := 0, len(bytes)-1

	for left < right {
		// 左指针找元音
		for left < right && !isVowel(bytes[left]) {
			left++
		}
		// 右指针找元音
		for left < right && !isVowel(bytes[right]) {
			right--
		}
		// 交换
		if left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}

	return string(bytes)
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
