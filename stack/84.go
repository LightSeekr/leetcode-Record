package stack

func largestRectangleArea(heights []int) int {

	n := len(heights)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}
	// 存下标
	stack := make([]int, 0)
	res := 0
	for idx, h := range heights {
		if len(stack) == 0 {
			stack = append(stack, idx)
		} else {
			for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
				// 出栈
				i := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				// 计算面积
				// 宽度 = 右边界(idx) - 左边界(stack顶) - 1
				width := idx
				if len(stack) > 0 {
					width = idx - stack[len(stack)-1] - 1
				}
				area := heights[i] * width
				if area > res {
					res = area
				}
			}
			stack = append(stack, idx)
		}
	}
	for len(stack) > 0 {
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 计算面积
		// 宽度 = 右边界(n) - 左边界(stack顶) - 1
		width := n
		if len(stack) > 0 {
			width = n - stack[len(stack)-1] - 1
		}
		area := heights[i] * width
		if area > res {
			res = area
		}
	}
	return res
}

func LargestRectangleArea(heights []int) int {
	return largestRectangleArea(heights)
}
