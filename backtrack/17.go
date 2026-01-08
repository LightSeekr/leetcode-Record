package backtrack

/*
*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/
func letterCombinations(digits string) []string {
	// 1. 特殊情况处理
	if len(digits) == 0 {
		return []string{}
	}

	var phoneMap = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var res = make([]string, 0)
	// index: 当前处理到 digits 中的第几个数字
	// path: 当前已经生成的字母组合
	var backtrack func(int, string)
	backtrack = func(index int, path string) {
		// 边界条件
		if index == len(digits) {
			res = append(res, path)
			return
		}
		// 获取当前数字对应的字母表
		digit := digits[index]
		letters := phoneMap[digit]

		// 遍历可选的字母
		for i := 0; i < len(letters); i++ {
			// 递归进入下一层
			// 这里直接传递 path + string(letters[i])
			// 因为字符串在 Go 中是不可变的，这相当于自动做了“撤销选择”的操作
			// 下一次循环时，path 依然是旧的那个 path
			backtrack(index+1, path+string(letters[i]))
		}
	}
	backtrack(0, "")
	return res
}
