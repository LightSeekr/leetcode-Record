package twoPointers

/*
*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。
*/
func maxArea(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	//如何优化这个双重循环呢
	res := 0
	// 先写一个暴力做法
	for idx, h := range height {
		for j := idx + 1; j < len(height); j++ {
			ans := (j - idx) * min(h, height[j])
			if ans > res {
				res = ans
				//fmt.Printf("idx:%d, j: %d, ans: %d\n", idx, j, ans)
			}
		}
	}
	return res

}

func maxArea1(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		hl, hr := height[l], height[r]
		area := (r - l) * min(hl, hr)
		res = max(res, area)
		if hl < hr {
			l++
		} else {
			r--
		}
	}
	return res
}
