package strings

func lengthOfLastWord(s string) int {
	n := len(s)
	start, end := -1, 0
	for i := n - 1; i >= 0; i-- {
		if s[i] != ' ' {
			end = i
			break
		}
	}

	for i := end; i >= 0; i-- {
		if s[i] == ' ' {
			start = i
			break
		}
	}

	var res string
	// 如果前面没有''
	if start == -1 {
		res = s[0 : end+1]
	} else {
		res = s[start+1 : end+1]
	}
	return len(res)
}
