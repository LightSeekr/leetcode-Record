package martix

import "fmt"

/*
*
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
第一步：转置 (对角线翻转)
第二步：翻转每一行 -> 得到顺时针旋转 90 度结果
线性代数的知识点。
*/
func rotate(matrix [][]int) {
	// 转置
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 翻转每一行
	for _, m := range matrix {
		for i, j := 0, len(m)-1; j > i; {
			m[i], m[j] = m[j], m[i]
			i++
			j--
		}
	}

}

func rotate180(matrix [][]int) {
	n := len(matrix)

	// 1. 上下翻转 (Flip Vertically)
	// 交换第 i 行和倒数第 i 行
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	// 2. 左右翻转 (Flip Horizontally)
	// 对每一行进行反转
	for i := 0; i < n; i++ {
		left, right := 0, n-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
	fmt.Println(matrix)
}
