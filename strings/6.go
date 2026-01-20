package strings

import "strings"

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	// 创建 numRows 个 Builder 来存每一行的字符
	rows := make([]strings.Builder, numRows)

	curRow := 0
	flag := -1 // 初始设为 -1，因为第一步我们会反转为 1 (向下)

	for _, char := range s {
		// 填入当前字符
		rows[curRow].WriteRune(char)

		// 判断是否需要转向
		// 到了第一行或者最后一行，都要转向
		if curRow == 0 || curRow == numRows-1 {
			flag = -flag
		}

		// 更新行号
		curRow += flag
	}

	// 拼接所有行
	var res strings.Builder
	for _, row := range rows {
		res.WriteString(row.String())
	}

	return res.String()
}
