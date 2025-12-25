package hash

func isHappy(n int) bool {
	numSet := make(map[int]struct{})
	for n != 1 {
		numSet[n] = struct{}{}
		n = changeNum(n)
		if _, ok := numSet[n]; ok {
			return false
		}
		if n == 1 {
			return true
		}
	}
	return true
}
func changeNum(n int) int {
	res := 0
	for n > 0 {
		num := n % 10
		res += num * num
		n /= 10
	}
	return res
}

func IsHappy(n int) bool {
	return isHappy(n)
}
