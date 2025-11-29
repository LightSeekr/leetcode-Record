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
