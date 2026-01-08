package backtrack

/*
*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。
如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，
其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
*/
func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	var backtrack func(int, int, int) bool
	backtrack = func(i, j, index int) bool {
		if index == len(word) {
			return true
		}
		if i < 0 || i >= row || j < 0 || j >= col || board[i][j] != word[index] {
			return false
		}
		// 这一格子匹配了，再走下一步
		found := false
		temp := board[i][j]
		board[i][j] = '#'

		found = backtrack(i+1, j, index+1) || backtrack(i-1, j, index+1) || backtrack(i, j+1, index+1) || backtrack(i, j-1, index+1)

		board[i][j] = temp

		return found
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] {
				// 可能不只有一个起点,所以每个起点都要遍历一遍
				// 这里找起点是为了剪枝条，避免无谓的遍历
				if backtrack(i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}

/**
我们需要在网格中找到一条路径，使得路径上的字符按顺序连起来等于给定的单词。

遍历起点：我们遍历网格中的每一个点。如果某个点 board[i][j] 等于单词的第一个字符 word[0]，那么就以这个点为起点开始 DFS 搜索。

DFS 搜索：
检查当前格子是否越界。
检查当前格子的字符是否匹配单词中对应的字符。
检查当前格子是否已经被访问过（避免走回头路）。

回溯 (Backtracking)：
为了标记“已访问”，我们可以把当前格子的字符临时修改为一个特殊字符（比如 '#'），或者使用一个 visited 数组。
如果发现这条路走不通（四个方向都找不到下一个字符），我们需要还原现场，把 '#' 改回原来的字符，让上层递归可以尝试别的路径。
*/
