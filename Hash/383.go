package Hash

func canConstruct(ransomNote string, magazine string) bool {
	// 长度预判：赎金字符串更长直接失败
	if len(ransomNote) > len(magazine) {
		return false
	}
	// 这里给一个容量 hint，避免 map 多次扩容（不是必须，只是小优化）
	strSet := make(map[rune]int, len(magazine))
	for _, n := range magazine {
		if _, ok := strSet[n]; ok {
			strSet[n] = strSet[n] + 1
		} else {
			strSet[n] = 1
		}
	}
	for _, n := range ransomNote {
		if _, ok := strSet[n]; ok {
			strSet[n] = strSet[n] - 1
			if strSet[n] < 0 {
				return false
			}
		} else {
			// 存在
			return false
		}
	}
	return true
}

func CanConstruct(ransomNote string, magazine string) bool {
	return canConstruct(ransomNote, magazine)
}

// canConstructArray 使用定长数组统计字符频次，
// 针对 LeetCode 383 只包含小写字母 a-z 的场景。 用字符串做哈希表
func canConstructArray(ransomNote, magazine string) bool {
	// 长度预判：赎金字符串比杂志长，肯定无法构造
	if len(ransomNote) > len(magazine) {
		return false
	}

	var cnt [26]int

	// 统计 magazine 中每个字符出现的次数
	for i := 0; i < len(magazine); i++ {
		cnt[magazine[i]-'a']++
	}

	// 消耗 ransomNote 中需要的字符
	for i := 0; i < len(ransomNote); i++ {
		idx := ransomNote[i] - 'a'
		cnt[idx]--
		if cnt[idx] < 0 {
			return false
		}
	}
	return true
}
