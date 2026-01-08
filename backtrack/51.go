package backtrack

/**
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
*/

func solveNQueens(n int) [][]string {
	var res [][]string

	// board[i] = j 表示第 i 行的皇后放在第 j 列
	// 这个数组只是为了最后生成输出结果用
	queens := make([]int, n)

	// 三个辅助布尔数组，用来快速判断冲突
	// 空间换时间，避免 O(N) 的遍历检查
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n) // 对应 row - col
	diag2 := make([]bool, 2*n) // 对应 row + col

	// 定义回溯函数
	var backtrack func(row int)
	backtrack = func(row int) {
		// 1. 结束条件：所有行都放好了
		if row == n {
			board := generateBoard(queens, n)
			res = append(res, board)
			return
		}

		// 2. 遍历当前行的每一列
		for col := 0; col < n; col++ {
			// 计算对角线索引
			// row - col 的范围是 [-(n-1), n-1]，加上 n 变成正数索引
			d1Index := row - col + n
			// row + col 的范围是 [0, 2n-2]
			d2Index := row + col

			// 3. 检查冲突
			if cols[col] || diag1[d1Index] || diag2[d2Index] {
				continue
			}

			// --- 做选择 ---
			queens[row] = col
			cols[col] = true
			diag1[d1Index] = true
			diag2[d2Index] = true

			// --- 递归下一行 ---
			backtrack(row + 1)

			// --- 撤销选择 (回溯) ---
			cols[col] = false
			diag1[d1Index] = false
			diag2[d2Index] = false
		}
	}

	backtrack(0)
	return res
}

// 辅助函数：将 queens 数组转换为题目要求的字符串格式
func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			if queens[i] == j {
				row[j] = 'Q'
			} else {
				row[j] = '.'
			}
		}
		board = append(board, string(row))
	}
	return board
}
