package strings

func longestCommonPrefix(strs []string) string {
	//1 <= strs.length <= 200
	// 0 <= strs[i].length <= 200
	res := ""
	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i] // 用第一个字符串中的每一个来做检查
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == 0 || i == len(strs[j]) || ch != strs[j][i] {
				// 出现空串
				// 第二个串短了
				//不符合情况
				return res
			}
		}
		res += string(ch)
	}
	return res
}
