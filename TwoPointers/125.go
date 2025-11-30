package TwoPointers

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	str := removeOtherChars(s)
	fmt.Println("str: ", str)
	for i, j := 0, len(str)-1; i < j; {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// removeOtherChars 将字符串中的大写字符转换为小写字符，并移除所有非字母数字字符
// 输入: 原始字符串s
// 输出: 处理后的字符串，只包含小写字母和数字
func removeOtherChars(s string) string {
	// 使用strings.Builder高效构建字符串
	var builder strings.Builder
	// 预先分配容量，避免多次扩容
	builder.Grow(len(s))

	// 遍历字符串中的每个字符
	for _, char := range s {
		// 检查字符是否为字母或数字
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			// 将字母转换为小写并添加到结果中
			builder.WriteRune(unicode.ToLower(char))
		}
		// 非字母数字字符将被自动忽略（移除）
	}
	// 返回处理后的字符串
	return builder.String()
}

func IsPalindrome(s string) bool {
	return isPalindrome(s)
}
