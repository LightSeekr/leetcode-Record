package Hash

import "strings"

func wordPattern(pattern string, s string) bool {
	//和之前的一样，只不过变个形式
	p2s := make(map[uint8]string)
	s2p := make(map[string]uint8)

	str := strings.Split(s, " ")
	if len(pattern) != len(str) {
		return false
	}
	length := len(pattern)
	for i := 0; i < length; i++ {
		pi := pattern[i]
		si := str[i]
		if s2, ok := p2s[pi]; ok && s2 != si {
			return false
		}
		if p2, ok := s2p[si]; ok && p2 != pi {
			return false
		}
		p2s[pi] = si
		s2p[si] = pi
	}
	return true

}

func WordPattern(pattern string, s string) bool {
	return wordPattern(pattern, s)
}

// WordPatternOptimized 是 wordPattern 函数的优化版本
// 优化点说明：
// 1. 预分配map容量：为两个映射表预分配容量，避免动态扩容
// 2. 使用range遍历：简化循环逻辑，提高代码可读性
// 3. 变量命名优化：使用更具描述性的变量名
// 4. 添加清晰注释：详细解释关键逻辑，提高代码可维护性
// 5. 移除冗余函数调用层：直接在主函数中实现逻辑
func WordPatternOptimized(pattern string, s string) bool {
	// 预分配map容量以减少动态扩容，提高性能
	p2s := make(map[uint8]string, len(pattern))
	s2p := make(map[string]uint8, len(pattern))

	// 拆分字符串并立即验证长度
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	// 使用range遍历pattern，简化索引管理
	for i, char := range pattern {
		word := words[i]

		// 检查字符到单词的映射是否一致
		if existingWord, exists := p2s[uint8(char)]; exists && existingWord != word {
			return false
		}

		// 检查单词到字符的映射是否一致
		if existingChar, exists := s2p[word]; exists && existingChar != uint8(char) {
			return false
		}

		// 更新双向映射关系
		p2s[uint8(char)] = word
		s2p[word] = uint8(char)
	}

	return true
}
