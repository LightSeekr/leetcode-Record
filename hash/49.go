package hash

func GroupAnagrams(strs []string) [][]string {
	return groupAnagrams(strs)
}

// GroupAnagrams groups together words that are anagrams of each other.
// 这里约定只包含小写字母 'a' ~ 'z'。
//
// 实现思路（针对小写字母优化）：
// - 用长度为 26 的整型数组记录每个字符的出现次数作为 key（[26]int 可以直接作为 map 的 key）
// - 相同 key 的字符串归为一组
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	// AI 写的比我想的好啊。
	/**
	首先 如何判断两个单词是字母异位词
	借助hash 记录单词中字母出现的数量，构建一个字符 到数字的映射，
		比如 eat 和 ate 中都是 a t e 各自出现了一次。所以二者都是 map[e]=1 map[a]=1 map[t]=1
	然后 比较两个map 就好了
	不过由于这里限定了是 26个小写字母，那么就可以更进一步，用数组来代替 map，然后比较两个数组是否相同
	只要对应的数组相等，那么两个单词就是相等的。
	那就可以用这个数组来做 key，将符合的字符串当作 value 这样就高效的完成了划分

	*/
	m := make(map[[26]int][]string, len(strs))
	for _, s := range strs {
		var key [26]int
		for i := 0; i < len(s); i++ {
			// 只考虑小写字母 'a' ~ 'z'
			key[s[i]-'a']++
		}
		m[key] = append(m[key], s)
	}

	res := make([][]string, 0, len(m))
	for _, group := range m {
		res = append(res, group)
	}
	return res
}
