package greedy

/*
*
步骤 1：预处理最远位置
首先，我们要知道每个字母在字符串中最后一次出现的位置。

遍历一遍字符串，用一个数组（或哈希表）lastPos 记录下来。
例如：lastPos['a'] = 8。
步骤 2：贪心扫描
我们需要确定一个片段的结束位置。

假设当前片段从 start 开始。
我们遍历这个片段里的每一个字符 c。
为了保证“同一字母最多出现在一个片段中”，如果我们把字符 c 划进了当前片段，那么这个片段的结束位置至少要延伸到 c 最后出现的位置 lastPos[c]。
我们维护一个变量 end，表示当前片段所需的最小结束位置。
end = max(end, lastPos[c])
当我们遍历到了 end 这个位置（i == end），说明：
在这个范围内所有出现过的字母，它们的最后一次出现位置都在 end 之前（或就是 end）。
我们找到了一个合法的切割点！
记录长度 end - start + 1，然后开启下一个片段（start = i + 1）。
*/
func partitionLabels(s string) []int {
	// 1. 记录每个字符最后出现的位置
	lastPos := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		lastPos[s[i]] = i
	}
	var res []int
	start := 0 // 当前片段的起始位置
	end := 0   // 当前片段的最远结束位置

	for i := 0; i < len(s); i++ {
		// 更新当前片段需要延伸到的最远位置
		end = max(end, lastPos[s[i]])
		// 如果遍历到了这个“最远边界”，说明可以切割了
		if i == end {
			res = append(res, end-start+1)
			start = i + 1 // 下一个片段的起点
		}
	}
	return res
}
