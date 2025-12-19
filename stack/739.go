package stack

/*
*
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，
其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。
如果气温在这之后都不会升高，请在该位置用 0 来代替。

[73,74,75,71,69,72,76,73]
[1,1,4,2,1,1,0,0]
*/
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	// 存下标
	stack := make([]int, 0)
	for idx, t := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, idx)
		} else {
			for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
				// 出栈
				i := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				res[i] = idx - i
			}
			stack = append(stack, idx)
		}
	}
	for len(stack) > 0 {
		i := stack[len(stack)-1]
		res[i] = 0
		stack = stack[:len(stack)-1]
	}
	return res
}

func DailyTemperatures(temperatures []int) []int {
	return dailyTemperatures(temperatures)
}
