package backtrack

/*
*
给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
*/
func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)

	var backtrack func(start int) // 从哪个字符串开始选择
	backtrack = func(start int) {
		// 确保选到的都是回文串
		if start == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := start; i < len(s); i++ {
			if checkStr(s, start, i) {
				// 选择 s[start,i] 这一部分
				path = append(path, s[start:i+1])
				backtrack(i + 1) //
				path = path[:len(path)-1]
			}

		}
	}

	backtrack(0)
	return res
}

func checkStr(str string, st, en int) bool {
	for i, j := st, en; i < j; {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

/**
先从s 中拿一个回文串，str_hui + str_next =s
下一步就是将 str_next 分割成一些子串，每一串都是回文串，递归了。
*/
