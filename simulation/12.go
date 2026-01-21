package simulation

import "strings"

/*
*
构建「数值 - 罗马字符」的降序映射列表（从大到小排列所有关键数值和对应字符）；
遍历该列表，对每个数值组合执行：
只要当前 num 大于等于该数值，就将对应的罗马字符追加到结果中，并从 num 中减去该数值；
重复上述操作，直到 num 小于该数值，再处理下一个更小的数值组合；
当 num 减为 0 时，结束遍历，结果即为最终罗马数字。
*/
func intToRoman(num int) string {
	// 1. 构建降序的数值-罗马字符映射列表（关键：必须从大到小）
	valueSymbols := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var res strings.Builder
	for _, symbol := range valueSymbols {
		for num >= symbol.value {
			res.WriteString(symbol.symbol)
			num -= symbol.value
		}
		if num == 0 {
			break
		}
	}
	return res.String()
}
