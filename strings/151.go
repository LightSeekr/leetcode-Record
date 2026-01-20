package strings

import "strings"

// 剪除空格
// "the sky is blue"
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
