package martix

/*
*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
*/
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	row, col := 0, n-1
	//
	for row < m && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			//每列的元素从上到下升序排列。所以 col 这一列都大于 target
			// 要缩小 matrix[row][col]
			col--
		} else {
			//每行的元素从左到右升序排列。 row 这一行都
			row++
		}
	}
	return false
}

/**
要么二维转一维，二分
要么利用这个性质：
每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

右上 或者左下角的元素可以作为判断标准，每次减少一行或者一列

*/
