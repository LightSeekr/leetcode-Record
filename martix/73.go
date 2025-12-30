package martix

/**
给定一个 m x n 的矩阵，
如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
*/

/*
*
标记出，哪些是原始的 0
哪些是之后被变成的 0
要求常量级别的空间，那还是算了。
可以使用二维转一维的做法，将原始的 0 的坐标保存下来。
最多需要
m := len(matrix)
n := len(matrix[0])
a[i][j] = i*n+j // 二维转一维度
一维转二维
i= temp/n
j =temp%n
*/
func setZeroes(matrix [][]int) {
	m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])
	if n == 0 {
		return
	}

	firstRowZero := false
	firstColZero := false

	// 1. 检查第一列是否需要置零
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	// 2. 检查第一行是否需要置零
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	// 3. 使用第一行/列作为标记，记录其他部分的零点
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 4. 根据标记将相应元素置零（不包括第一行/列）
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 5. 如果需要，将第一行置零
	if firstRowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// 6. 如果需要，将第一列置零
	if firstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
