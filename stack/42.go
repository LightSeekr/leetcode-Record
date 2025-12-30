package stack

func trap(height []int) (ans int) {
	st := []int{}
	for i, h := range height {
		for len(st) > 0 && height[st[len(st)-1]] <= h {
			bottomH := height[st[len(st)-1]]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			left := st[len(st)-1]
			dh := min(height[left], h) - bottomH // 面积的高
			ans += dh * (i - left - 1)
		}
		st = append(st, i)
	}
	return
}
