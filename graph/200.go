package graph

// 泛洪算法
func numIslandsDfs(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, x, y int) {
	n, m := len(grid), len(grid[0])
	if x >= n || x < 0 || y >= m || y < 0 || grid[x][y] == '0' {
		return
	}
	// 原地修改切片，保证
	grid[x][y] = '0'
	dfs(grid, x, y-1)
	dfs(grid, x, y+1)
	dfs(grid, x+1, y)
	dfs(grid, x-1, y)
}

// bfs 做法  注意入队的时候就标记为 0 防止重复入队。而不是出队的时候标记为0
func numIslandsBfs(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	res := 0
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := make([][]int, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				res++
				grid[i][j] = '0'
				queue = append(queue, []int{i, j})
				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:]
					for _, d := range direction {
						x := cur[0] + d[0]
						y := cur[1] + d[1]
						if x >= n || x < 0 || y >= m || y < 0 || grid[x][y] == '0' {
							continue
						}
						// 合理且为 1 的话
						grid[x][y] = '0' // 入队时标记，防止重复入队
						queue = append(queue, []int{x, y})
					}
				}
			}
		}
	}
	return res
}
