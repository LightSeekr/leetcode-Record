package Hash

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	strMap := make(map[rune]int)
	for _, v := range s {
		strMap[v]++
	}
	for _, v := range t {
		if cnt := strMap[v]; cnt == 0 {
			return false
		} else {
			strMap[v]--
		}
	}
	return true
}

func IsAnagram(s string, t string) bool {
	return isAnagram(s, t)
}
