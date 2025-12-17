package BinarySearch

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// 边界case 处理
	if matrix[m-1][n-1] < target || matrix[0][0] > target {
		return false
	}

	// 题目中有 m,n>=1
	//if m == 0 || n == 0 { }
	l, r := -1, (m-1)*n+n
	for r-l > 1 {
		mid := l + (r-l)/2
		// 一维转2维
		if matrix[mid/n][mid%n] >= target {
			r = mid
		} else {
			l = mid
		}
	}

	// r 是 >= target 的左边界
	// l是  <target 的右边界

	// 如果数组全都比 target 小呢 要么在这里处理，要么在上面处理
	if r == (m-1)*n+n {
		return false
	}

	if matrix[r/n][r%n] == target {
		return true
	} else {
		return false
	}
}

func searchMatrix1(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	i, j := 0, m-1
	for i < n && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}

	return false
}

func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, target)
}
