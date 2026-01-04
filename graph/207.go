package graph

// 拓扑排序，判断图是否有环
// 图怎么存？会不会写？

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	/// inDegree 存储 课程i 需要的先修课程数量，也就是 入度
	inDegree := make([]int, numCourses)
	for _, r := range prerequisites {
		//prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
		cur := r[0]
		pre := r[1]
		// 构建一条从 bi 到 ai 的有向边
		adj[pre] = append(adj[pre], cur)
		inDegree[cur]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	finishCount := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		finishCount++

		for _, next := range adj[cur] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	return finishCount == numCourses
}
