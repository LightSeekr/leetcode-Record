package TwoPointers

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	// 判断 s 是否是 t 的子序列
	//  两个指针，不改变相对位置，那么就可以使用 双指针
	idx1, idx2 := 0, 0
	for ; idx1 < len(s) && idx2 < len(t); idx2++ {
		if s[idx1] == t[idx2] {
			idx1++
		}
	}
	return idx1 == len(s)
}

func IsSubsequence(s string, t string) bool {
	return isSubsequence(s, t)
}
