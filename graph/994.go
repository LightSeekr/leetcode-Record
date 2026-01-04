package graph

// fresh 的运用，以看出bfs后还有没有新鲜橘子，还能避免最后一遍扫描队列
// 提醒大家认真读题，初始化就有腐烂的橘子，还是多个
func orangesRotting(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	res := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := make([][]int, 0)
	fresh := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	for fresh > 0 && len(queue) > 0 {
		res++ // 经过一分钟
		length := len(queue)
		for k := 0; k < length; k++ {
			// 出队列
			cur := queue[0]
			queue = queue[1:]
			for _, direction := range directions {
				x := cur[0] + direction[0]
				y := cur[1] + direction[1]
				if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] == 1 {
					grid[x][y] = 2 // 变成腐烂橘子
					queue = append(queue, []int{x, y})
					fresh--
				}
			}
		}
	}

	if fresh > 0 {
		return -1
	}
	return res
}
