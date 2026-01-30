package slideWindow

import "fmt"

/*
*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
*
*/

func lengthOfLongestSubstring3(s string) int {
	// 滑动窗口
	// r++ 将字符都放进去。
	// 如果出现重复，窗口移动,l++
	// 更新长度
	Map := make(map[byte]int)
	res, l := 0, 0
	for r := 0; r < len(s); r++ {
		ch := s[r]
		if idx, ok := Map[ch]; ok {
			// 当前字符存在重复
			l = max(l, idx+1)
		}
		Map[ch] = r
		res = max(res, r-l+1)
	}
	return res
}

func lengthOfLongestSubstring2(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	// 这题要求的是不重复的子串，每个字符最多只能出现一次，不需要计数，所以用集合
	chMap := make(map[rune]bool)
	l, ans := 0, 0

	for r, ch := range s {
		// 枚举右边，准备把 ch 放入窗口
		// ch 存在，说明窗口内出现了重复字符，这时候应该收缩窗口了
		for chMap[ch] {
			chMap[rune(s[l])] = false
			l++
		}
		// 更新结果
		ans = max(ans, r-l+1)
		// 再把 r 放入其中
		chMap[rune(s[r])] = true
	}
	return ans
}

// 没问题，就是写的啰嗦了。
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	// 用map 来判断重复, true 表示加入，false 表示移出或者不存在
	charSet := make(map[rune]bool)
	l, ans := 0, -1
	for r, ch := range s {
		fmt.Printf("r=%d\n", r)
		// s[r] 所在位置的元素重复了
		if v := charSet[ch]; v {
			// 此时[l,r-1] 这个子串符合条件
			ans = max(ans, r-l) // r-1-l+1
			fmt.Printf("charSet[%c]重复,[l,r-1] 这个子串符合条件\n l: %d,r: %d,ans: %d str:%s\n", ch, l, r, ans, s[l:r])
			// 左移动 l
			for {
				charSet[rune(s[l])] = false
				l++
				if !charSet[ch] {
					break
				}
			}
			fmt.Printf("移动l后,l: %d,r: %d\n", l, r)
			// 出来了之后 此时的 [l,r] 符合条件
			ans = max(ans, r-l+1)
			// 要把r 放入窗口
			charSet[ch] = true
			fmt.Println("charSet: ", charSet)
		} else {
			// 不存在重复直接放入
			charSet[ch] = true
			fmt.Printf("%c 不存在,直接放入map \n", ch)
			fmt.Println("map:", charSet)
		}
		fmt.Println("=====================================")
	}
	//l,r 如果一直没有重复呢
	ans = max(ans, n-l)
	return ans
}

func LengthOfLongestSubstring(s string) int {
	//return lengthOfLongestSubstring(s)
	return lengthOfLongestSubstring2(s)

}
