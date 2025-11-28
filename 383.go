package leetcode

func CanConstruct(ransomNote string, magazine string) bool {
	m1 := make(map[rune]int, 1024)

	for _, r := range magazine {
		if val, ok := m1[r]; ok {
			m1[r] = val + 1
		} else {
			m1[r] = 1
		}
	}

	//fmt.Println(m1)

	for _, r := range ransomNote {
		if val, ok := m1[r]; ok {
			val = val - 1
			if val < 0 {
				return false
			}
			m1[r] = val
		} else {
			return false
		}
	}

	return true
}
