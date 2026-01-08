package dynamicProgramming

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	if numRows == 0 {
		return res
	}
	if numRows >= 1 {
		res = append(res, []int{1})
	}
	if numRows >= 2 {
		res = append(res, []int{1, 1})
	}

	// 第三行开始，下标为 2
	for i := 2; i < numRows; i++ {
		// 每一行有 i 个数
		temp := make([]int, i+1)
		temp[0] = 1
		temp[i] = 1
		for j := 1; j < i; j++ {
			temp[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, temp)
	}

	return res
}
