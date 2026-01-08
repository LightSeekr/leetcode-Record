package backtrack

func generateParenthesis(n int) []string {
	var res []string
	var backtrack func(string, int, int)

	// 提前剪枝
	backtrack = func(path string, open int, close int) {
		if len(path) == 2*n {
			res = append(res, path)
			return
		}

		// 2. 尝试添加左括号
		// 条件：左括号还没用完
		if open < n {
			backtrack(path+"(", open+1, close)
		}

		// 3. 尝试添加右括号
		// 条件：右括号的数量必须小于左括号（意味着有左括号还没闭合）
		if close < open {
			backtrack(path+")", open, close+1)
		}

	}

	backtrack("", 0, 0)
	return res
}
