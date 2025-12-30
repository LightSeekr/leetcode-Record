package martix

/**
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
设定每次循环的边界
拐弯的边界
技巧
dx dy
判断转弯
*/

func spiralOrder(matrix [][]int) []int {
	// 右 下 左 上
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	d := 0
	x, y := 0, 0

	m, n := len(matrix), len(matrix[0])
	isvisited := make([][]bool, m)
	for i := range isvisited {
		isvisited[i] = make([]bool, n)
	}

	res := make([]int, 0)
	for i := 0; i < m*n; i++ {
		res = append(res, matrix[x][y])
		isvisited[x][y] = true
		a := x + dx[d]
		b := y + dy[d]
		if a < 0 || a >= m || b < 0 || b >= n || isvisited[a][b] {
			d = (d + 1) % 4
			a = x + dx[d]
			b = y + dy[d]
		}
		x, y = a, b
	}
	return res
}
