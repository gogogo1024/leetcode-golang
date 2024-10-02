package slidewindow

// 给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。
// 子串 (Substring) 是指字符串中连续的字符组成的片段。
// 例如, 对于字符串"abcxyz", 它的子串有"bcx"、"xyz" 等.
// 子串必须是连续的, 子序列可以是不连续的。
// 子串是字符串的一部分, 子序列不必是字符串的一部分。
// 任何子串都可以看成是一个子序列, 但是子序列不一定是子串。
// 1
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是"abc"，所以其长度为 3。
// 2
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是"b"，所以其长度为 1。
// 3 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
// 自定义约束接口，表示可以进行比较的类型
type Ordered interface {
	~int | ~float64 | ~string
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	n := len(s)
	right, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for right+1 < n && m[s[right+1]] == 0 {
			m[s[right+1]]++
			right++
		}
		ans = max(ans, right-i+1)
	}
	return ans
}

// 假设你有两个数组，一个长一个短，短的元素均不相同。找到长数组中包含短数组所有
// 的元素的最短子数组，其出现顺序无关紧要。返回最短子数组的左端点和右端点，如有
// 多个满足条件的子数组，返回左端点最小的一个。若不存在，返回空数组。
// 1
// 输入:
// big = [7,5,9,0,2,1,3,5,7,9,1,1,5,8,8,9,7]
// small = [1,5,9] 输出: [7,10]
// 2
// 输入:
// big = [1,2,3]
// small = [4] 输出: []
func shortestSeq[T Ordered](big []T, small []T) []int {
	bigLen := len(big)
	smallLen := len(small)
	if smallLen > bigLen {
		return []int{}
	}

	// 结果区间
	ans := []int{}
	minLen := bigLen + 1

	// 统计 small 数组中每个元素的个数
	require := make(map[T]int)
	for _, e := range small {
		require[e]++
	}

	// 记录窗口中需要匹配的元素数
	matchCount := 0
	window := make(map[T]int)
	l := 0

	for r := 0; r < bigLen; r++ {
		// 当前元素
		rightElem := big[r]

		// 如果当前元素在 small 中，更新窗口的统计信息
		if _, ok := require[rightElem]; ok {
			window[rightElem]++

			// 窗口中的该元素数量满足 small 中要求时，增加匹配数
			if window[rightElem] == require[rightElem] {
				matchCount++
			}
		}

		// 当所有元素匹配时，尝试收缩窗口
		for matchCount == len(require) {
			if r-l+1 < minLen {
				minLen = r - l + 1
				ans = []int{l, r}
			}

			// 准备收缩窗口，移除左边元素
			leftElem := big[l]
			if _, ok := require[leftElem]; ok {
				// 窗口中的该元素数量减少
				if window[leftElem] == require[leftElem] {
					matchCount--
				}
				window[leftElem]--
			}
			l++
		}
	}
	return ans
}
