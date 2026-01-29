package strings

import "strings"

/**
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。


*/
/**
 1. 先去除首尾空格
 2. 按空格分割字符串得到单词数组
 3. 从后往前遍历单词数组，跳过空字符串（多个空格导致）
 4. 使用 StringBuilder 拼接结果，最后去除尾部空格

 时间复杂度：O(n)，空间复杂度：O(n)
**/
func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	split := strings.Split(s, " ")
	res := strings.Builder{}
	for j := len(split) - 1; j >= 0; j-- {
		if len(split[j]) == 0 {
			continue
		}
		res.WriteString(split[j])
		res.WriteRune(' ')
	}
	// 剪除右边空格
	return strings.TrimSpace(res.String())
}

/*
*

	双指针法：从后往前遍历，直接定位每个单词的边界
	1. 使用 right 指针跳过尾部空格，定位单词末尾
	2. 使用 left 指针向左移动，定位单词开头
	3. 提取单词并添加到结果中

	优点：避免 Split 产生的中间数组，无需处理空字符串
	时间复杂度：O(n)，空间复杂度：O(n)

*
*/
func reverseWords2(s string) string {
	res := strings.Builder{}
	right := len(s) - 1

	for right >= 0 {
		// 跳过尾部空格
		for right >= 0 && s[right] == ' ' {
			right--
		}
		if right < 0 {
			break
		}

		// 定位单词开头
		left := right
		for left >= 0 && s[left] != ' ' {
			left--
		}

		// 添加单词
		if res.Len() > 0 {
			res.WriteByte(' ')
		}
		res.WriteString(s[left+1 : right+1])

		// 移动到下一个单词
		right = left
	}

	return res.String()
}
