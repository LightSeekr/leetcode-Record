package enumerate

import (
	"fmt"
)

func maxDistance(arrays [][]int) int {
	res := 0
	mn, mx := arrays[0][0], arrays[0][len(arrays[0])-1]

	// 维护左边所有数组中的最小值，最大值
	// for range 还可以这么写
	for _, arr := range arrays[1:] {
		nn, nx := arr[0], arr[len(arr)-1]
		res = max(res, abs(mx-nn))
		res = max(res, abs(nx-mn))
		fmt.Printf("mn = %d, mx = %d,nn=%d,nx=%d, res:=%d\n,", mn, mx, nn, nx, res)
		mn = min(mn, nn)
		mx = max(mx, nx)

	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxDistance2(arrays [][]int) int {
	res := 0
	// 维护啥呢？
	for i, arr := range arrays {
		for j := 0; j < i; j++ {
			a1, a2 := arr[0], arr[len(arr)-1]
			b1, b2 := arrays[j][0], arrays[j][len(arrays[j])-1]
			res = max(res, abs(b2-a1))
			res = max(res, abs(b1-a2))
		}
	}
	return res
}

func MaxDistance(arrays [][]int) int {
	return maxDistance2(arrays)
}
