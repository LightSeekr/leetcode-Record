package simulation

/*
*
先建立「字符 - 数值」的映射表，方便快速查询单个字符的值；
遍历字符串，比较当前字符值与下一个字符值：
若当前值 < 下一个值：说明是减法组合，结果 -= 当前值；
若当前值 ≥ 下一个值：正常累加，结果 += 当前值；
最后将最后一个字符的值加入结果（因为遍历到倒数第二个字符时，最后一个字符未处理）。
*/
func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)
	// 2. 遍历字符串（仅到倒数第二个字符）
	for i := 0; i < n-1; i++ {
		current := romanMap[s[i]]
		next := romanMap[s[i+1]]
		// 3. 判断当前值与下一个值的大小关系
		if current < next {
			total -= current // 减法组合：减去当前值
		} else {
			total += current // 正常情况：累加当前值
		}
	}
	// 4. 加上最后一个字符的值（最后一个字符无后续，直接累加）
	total += romanMap[s[n-1]]

	return total
}
