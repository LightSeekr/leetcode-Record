package slideWindow

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	ns := len(s)
	np := len(p)
	// 边界 case 处理
	if ns < np {
		return res
	}
	/**
	这题 p 的长度是固定的，那么意味着是否 窗口的大小也是固定的。
	窗口一滑动，
	如何进行一个计数器的状态统计。
	有几个字符，出现了几次？
	如何对计数器进行比较然后拿到结果？
	如何在不遍历 map 的情况下，进行一个字符串异位词的比较。
	1. 用数组代替 map。数组可以直接用原生的比较
	2.
	*/
	cp := [26]int{}
	cs := [26]int{}

	for i := 0; i < len(p); i++ {
		cp[p[i]-'a']++
	}
	l := 0
	//预处理
	//for r := 0; r < np-1; r++ {
	//	cs[s[r]-'a']++
	//}

	// 如何 r 从 np-1 开始的话，就需要做预处理
	for r := 0; r < ns; r++ {
		// 应该在这里，优化，不用遍历map 就可以得到答案
		cs[s[r]-'a']++
		l = r - np + 1
		if l < 0 {
			continue
		}
		if cs == cp {
			res = append(res, l)
		}
		cs[s[l]-'a']-- // l左边移动
	}
	return res
}

func findAnagrams2(s string, p string) []int {
	ns, np := len(s), len(p)
	res := make([]int, 0)
	// 找 s 中所有 p 的异味词的子串
	if ns < np {
		return res
	}
	cnts, cntp := [26]int{}, [26]int{}

	for _, ch := range p {
		cntp[ch-'a']++
	}

	for r := 0; r < ns; r++ {
		// 这里还是把元素一个个放入窗口
		cnts[s[r]-'a']++
		l := r - np + 1 // p 的长度固定，枚举 r的时候，计算l 实际上是计算窗口内元素是否足够
		if l < 0 {
			// 说明窗口中的元素不足，还要继续放
			continue
		}
		// 当窗口中的元素足够了，
		if cnts == cntp {
			// 说明找到了异味词,向结果中加入 l
			res = append(res, l)
		}
		//删除元素
		cnts[s[l]-'a']--
	}
	return res
}

//1225
// 没用滑窗，直接用数据结构锤开了。
func findAnagrams3(s string, p string) []int {
	lens, lenp := len(s), len(p)
	//if lens < lenp {
	//	return []int{}
	//}
	// 用数组做哈希表的 key，值为[]切片的 int 记录开始的位置
	Map := make(map[[26]int][]int)
	// 从i+lenp <= lens; 这个条件就可以发现上面那个 if 可以少写一个
	for i := 0; i+lenp <= lens; i++ {
		cnt := [26]int{}
		// 用数组做哈希表
		for _, r := range s[i : i+lenp] {
			cnt[r-'a'] += 1
		}
		// 构建 value
		if _, ok := Map[cnt]; !ok {
			Map[cnt] = make([]int, 0)
		}
		Map[cnt] = append(Map[cnt], i)
	}
	cnt := [26]int{}
	for _, r := range p {
		cnt[r-'a'] += 1
	}
	return Map[cnt]
}
