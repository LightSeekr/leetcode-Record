package hash

func isIsomorphic(s string, t string) bool {
	// 边界 case : len(s) == len(t)
	//  建立双向映射，看是否是空的
	s2t := make(map[rune]rune)
	t2s := make(map[rune]rune)
	for i := 0; i < len(s); i++ {
		a := rune(s[i])
		b := rune(t[i])
		if val, ok := s2t[a]; ok && val != b {
			return false
		}
		if val, ok := t2s[b]; ok && val != a {
			return false
		}
		s2t[a] = b
		t2s[b] = a
	}
	return true

}

func IsIsomorphic(s string, t string) bool {
	return isIsomorphic(s, t)
}
