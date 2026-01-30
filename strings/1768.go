package strings

/*
*
给你两个字符串 word1 和 word2 。请你从 word1 开始，通过交替添加字母来合并字符串。如果一个字符串比另一个字符串长，就将多出来的字母追加到合并后字符串的末尾。

返回 合并后的字符串 。
*/
/**
 1. 使用双指针 i 和 j 分别遍历两个字符串
  2. 交替添加字符到结果中，直到其中一个字符串遍历完
  3. 将较长字符串剩余的部分追加到结果末尾

  时间复杂度：O(n1 + n2)，空间复杂度：O(n1 + n2)
**/
func mergeAlternately(word1 string, word2 string) string {
	n1, n2 := len(word1), len(word2)
	result := make([]byte, 0, n1+n2)

	i, j := 0, 0
	for i < n1 && j < n2 {
		result = append(result, word1[i])
		result = append(result, word2[j])
		i++
		j++
	}

	// 追加剩余字符
	for i < n1 {
		result = append(result, word1[i])
		i++
	}
	for j < n2 {
		result = append(result, word2[j])
		j++
	}

	return string(result)
}
