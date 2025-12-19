package stack

import (
	"strings"
	"unicode"
)

/*
*
s 由小写英文字母、数字和方括号 '[]' 组成
s 中所有整数的取值范围为 [1, 300]
*/
func decodeString(s string) string {
	var numStack []int
	var strStack []string
	num := 0
	res := ""

	for _, char := range s {
		if unicode.IsDigit(char) {
			num = num*10 + int(char-'0')
		} else if char == '[' {
			numStack = append(numStack, num)
			strStack = append(strStack, res)
			num = 0
			res = ""
		} else if char == ']' {
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			res = prevStr + strings.Repeat(res, count)
		} else {
			res += string(char)
		}
	}
	return res
}

func DecodeString(s string) string {
	return decodeString(s)
}
