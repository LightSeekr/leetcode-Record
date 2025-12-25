package twoPointers

import (
	"fmt"
)

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
[0,1,0,2,1,0,1,3,2,1,2,1]

目前的这个解法是有问题的，他只适合于 计算右边比左边高的情况下的雨水
如果左边比右边高。那么就计算不了 [4,2,3]就是例子。
*/
func trap(height []int) (ans int) {
	n := len(height)
	// 记录下，右边第一个比他大的数的下标
	stack := make([]int, 0)
	right := make([]int, n)
	// 直接填充所有元素为 -1
	fillSlice(right, -1)
	//right[i] 就是右边第一个比 height[i] 大的下标
	for i := 0; i < n; i++ {
		//if len(stack) == 0 {
		//	stack = append(stack, i)
		//} else {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			// 出栈，记录答案
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			right[top] = i
			//fmt.Printf("right[%d]=%d\n", top, i)
		}
		stack = append(stack, i)
		//}
	}
	//fmt.Println(height)
	//fmt.Println("===========================")
	//fmt.Println(right)
	stack = make([]int, 0)
	left := make([]int, n)
	fillSlice(left, -1)
	//left[i] 就是左边第一个比 height[i] 大的下标
	for i := 0; i < n; i++ {
		//if len(stack) == 0 {
		//	stack = append(stack, i)
		//} else {
		//	for len(stack) > 0 && height[i] >= height[stack[len(stack)-1]] {
		//		// 直接出栈
		//		stack = stack[:len(stack)-1]
		//	}
		//	if len(stack) > 0 {
		//		left[i] = stack[len(stack)-1]
		//	}
		//	stack = append(stack, i)
		//}
		for len(stack) > 0 && height[i] >= height[stack[len(stack)-1]] {
			// 直接出栈
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	fmt.Println(height)
	fmt.Println(right)
	fmt.Println(left)
	fmt.Println("===========================")

	ans = 0
	for i := 0; i < n; i++ {
		if left[i] == -1 || right[i] == -1 {
			continue
		}
		ans += 1*min(height[left[i]], height[right[i]]) - height[i]
		fmt.Printf("i:%d left:%d right:%d ans:%d\n", i, left[i], right[i], ans)
	}
	//for i := 0; i < n; {
	//	// 先写， 写完了就知道该怎么优化了。
	//	//if i == 0 && height[0] == 0 {
	//	//	continue
	//	//} else {
	//	//	end := right[i]
	//	//	for st := i + 1; st < end; st++ {
	//	//		ans += (height[i] - height[st]) * 1
	//	//		fmt.Printf("i=%d,st=%d,height[i]=%d, height[st]=%d,ans=%d\n", i, st, height[i], height[st], ans)
	//	//
	//	//	}
	//	//}
	//	//这种算法有问题。 只适合于 计算右边比左边高的情况下的雨水
	//	//如果左边比右边高。那么就计算不了， [4,2,3]就是例子。
	//	end := right[i]
	//	//fmt.Printf("i=%d,end=%d\n", i, end)
	//	for st := i + 1; st < end; st++ {
	//		ans += (height[i] - height[st]) * 1
	//		fmt.Printf("height[%d]=%d, height[%d]=%d,ans=%d\n", i, height[i], st, height[st], ans)
	//	}
	//	//fmt.Println("===========================")
	//	// 更新 i 避免重复计算面积
	//	//i = right[i] //这么更新会让循环出不去。因为不存在的就是 0
	//	if end != -1 {
	//		i = end
	//	} else {
	//		i++
	//	}
	//}
	return ans
}

// trap1 是 trap 的正确实现版本：
// 用双指针维护左右两侧的最高柱（leftMax/rightMax），并在较低一侧结算该位置能存的水量。
// 接雨水位置 i 的水位应该由 左侧最高值 maxLeft[i] 和 右侧最高值 maxRight[i] 决定，即 min(maxLeft[i], maxRight[i]) -
//
//	height[i]；而你这里算的是 “最近的更高柱”（previous/next greater element），它不等价于 maxLeft/maxRight，会在“最近更
//	高柱不够高，但更远处有更高柱”时 少算。
func trap1(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	l, r := 0, n-1
	leftMax, rightMax := 0, 0
	ans := 0

	for l < r {
		if height[l] < height[r] {
			if height[l] >= leftMax {
				leftMax = height[l]
			} else {
				ans += leftMax - height[l]
			}
			l++
		} else {
			if height[r] >= rightMax {
				rightMax = height[r]
			} else {
				ans += rightMax - height[r]
			}
			r--
		}
	}
	return ans
}

// trap2 使用单调栈（存下标，栈内高度单调递减）实现。
// 当遇到比栈顶更高的柱子时，弹出“凹槽底”并用当前柱(i)与新的栈顶(left)作为左右边界结算面积。
func trap2(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	stack := make([]int, 0, n) // 栈里存下标
	ans := 0
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break // 左边没墙了
			}
			left := stack[len(stack)-1]
			width := i - left - 1
			boundedHeight := min(height[left], height[i]) - height[mid]
			if boundedHeight > 0 {
				ans += width * boundedHeight
			}
		}
		stack = append(stack, i)
		fmt.Println(stack)
		//printStack(stack, height)
	}

	return ans
}

// trap3 把 trap2 的“弹栈结算”思路，改写为：
// 1) 先预处理每个位置的 nextGreater（右侧第一个严格更高的柱子下标）
//2) 再预处理每个位置的 prevGreaterOrEqual（左侧第一个高度 >= 当前的柱子下标）
// 3) 最后按“弹栈时的 mid 结算公式”汇总面积
//
// 注意：这里不能像 trap 那样对每个 i 直接用 left/right 去算 1 列水量，
// 而是按 mid（凹槽底）去算 width * boundedHeight，这样才能避免漏算/错算。
//必须要一边大于一边大于等于。如果全都是大于呢？会发生什么？这改一下，用错误数据来回怼你就好了

func trap3(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	nextGreater := make([]int, n)
	fillSlice(nextGreater, -1)
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGreater[mid] = i
		}
		stack = append(stack, i)
	}

	prevGreaterOrEqual := make([]int, n)
	fillSlice(prevGreaterOrEqual, -1)
	stack = stack[:0]
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			prevGreaterOrEqual[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	ans := 0
	for mid := 0; mid < n; mid++ {
		left := prevGreaterOrEqual[mid]
		right := nextGreater[mid]
		if left == -1 || right == -1 {
			continue
		}
		width := right - left - 1
		boundedHeight := min(height[left], height[right]) - height[mid]
		if boundedHeight > 0 {
			ans += width * boundedHeight
		}
	}
	return ans
}
func printStack(stack []int, height []int) {
	for _, i := range stack {
		fmt.Printf("%d\t", height[i])
	}
	fmt.Println()
}
func fillSlice(res []int, val int) {
	for i := 0; i < len(res); i++ {
		res[i] = val
	}
}
