package slideWindow

import "fmt"

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

func lengthOfLongestSubstring(s string) int {
	// 这样写也没问题。
	// 注意是用hashMap 计数器，不仅判断是否存在，还要判断个数
	// 还是用 hashSet 集合，只需要判断存在即可
	// 现在的问题是，应该用循环还是用 if +循环呢？来判断结果 答案在什么时候可以更新？ labuladong 用的是 for
	n := len(s)
	if n <= 1 {
		return n
	}
	charMap := make(map[rune]int)
	l, ans := 0, 0
	for r, ch := range s {
		for charMap[ch] > 0 {
			//出现重复 移动 l
			// 这里 l 在之前判断过，所以这里一定存在。
			charMap[rune(s[l])] = charMap[rune(s[l])] - 1
			l++
		}
		ans = max(ans, r-l+1)
		charMap[ch] = 1
	}
	return ans
}

func lengthOfLongestSubstring_slide(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	charMap := make(map[rune]int)
	l, ans := 0, 0
	for r, ch := range s {
		//这里还有个问题就是 写 if 还是写 for ？
		// 写 for
		for charMap[ch] > 0 {
			//出现重复 移动 l
			// 这里 l 在之前判断过，所以这里一定存在。
			charMap[rune(s[l])] = charMap[rune(s[l])] - 1
			l++
		}
		ans = max(ans, r-l+1)
		charMap[ch] = 1
	}
	return ans
}

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
func lengthOfLongestSubstring1(s string) int {
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
