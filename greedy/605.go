package greedy

/**
假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。
*/

/*
*

 1. 贪心算法：遍历花坛，遇到能种花的位置就种

 2. 能种花的条件：当前位置为0，且左右相邻位置都为0（或越界


 3. 统计能种的花数，判断是否 >= n
    时间复杂度：O(m)，空间复杂度：O(1)

*
*/
 2. 能种花的条件：当前位置为0，且左右相邻位置都为0（或越界）
 3. 统计能种的花数，判断是否 >= n

 时间复杂度：O(m)，空间复杂度：O(1)
**/
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	length := len(flowerbed)

	for i := 0; i < length; i++ {
		// 当前位置为空
		if flowerbed[i] == 0 {
			// 检查左边（越界视为0）
			leftEmpty := i == 0 || flowerbed[i-1] == 0
			// 检查右边（越界视为0）
			rightEmpty := i == length-1 || flowerbed[i+1] == 0

			if leftEmpty && rightEmpty {
				flowerbed[i] = 1 // 种花
				count++
				if count >= n {
					return true
				}
			}
		}
	}

	return count >= n
}
