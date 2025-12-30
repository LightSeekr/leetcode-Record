package stack

import "fmt"

/*
*
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，
其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。
如果气温在这之后都不会升高，请在该位置用 0 来代替。

[73,74,75,71,69,72,76,73]
[1,1,4,2,1,1,0,0]
*/

// 从右往左遍历
func dailyTemperatures2(t []int) []int {
	n := len(t)
	res := make([]int, n)
	// 存下标
	stack := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		// 对于每一个 t[i],它的答案就在栈中，
		// 因为从右往左遍历，找到右边的第一个比其大的数
		if i == n-1 {
			// 这个是边界 case
			res[i] = 0
		} else {
			// 入栈之前的条件是什么？
			// 找当前元素的答案
			// 从右往左遍历，t[i] 右边的数进栈了，
			// if t[i]> t[stack[len(stack)-1]] 就说明 stack[len(stack)-1] 不可能成为答案了。
			// 因为 第一 t[i] 更大，第二 t[i] 更靠左，更能
			for len(stack) > 0 && t[i] >= t[stack[len(stack)-1]] {
				// 出栈
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				res[i] = 0
			} else {
				res[i] = stack[len(stack)-1] - i // 差值
			}
		}
		// 入栈
		stack = append(stack, i)
		printStack(stack, t)
	}
	return res
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	// 存下标
	stack := make([]int, 0)
	for idx, t := range temperatures {

		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			// 出栈
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res[i] = idx - i
		}
		stack = append(stack, idx)
		printStack(stack, temperatures)
	}

	for len(stack) > 0 {
		i := stack[len(stack)-1]
		res[i] = 0
		stack = stack[:len(stack)-1]
	}
	return res
}

func printStack(s []int, temperatures []int) {
	for _, t := range s {
		fmt.Printf("%d ", temperatures[t])
	}
	fmt.Println()
}
