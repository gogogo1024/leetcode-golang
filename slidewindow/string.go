package slidewindow

import "math"

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
				// 窗口中的该元素数量减少,因为只有同一个元素对应个数相等，l需要右移，就必须
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

// 给定两个字符串 s 和 t 。返回 s 中包含 t 的所有字符的最短子字符串。如果 s 中不存
// 在符合条件的子字符串，则返回空字符串"" 。如果 s 中存在多个符合条件的子字符串,
// 返回任意一个。
// 注意：对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 输入: s = "ADOBECODEBANC", t = "ABC"
// 输出: "BANC"
// 解释: "BANC" 包含了字符串 t 的所有字符'A','B','C'
func minWindow(s string, t string) string {
	// 边界条件
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// 记录 t 中的字符及其数量
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	// 用于滑动窗口中记录字符数量
	window := make(map[byte]int)

	// 左右指针
	left, right := 0, 0
	// 记录最小覆盖子串的起始位置和长度
	start, length := 0, math.MaxInt32
	// 记录窗口中满足条件的字符数
	valid := 0

	for right < len(s) {
		// 扩大窗口
		c := s[right]
		right++

		// 如果该字符在 t 中
		if need[c] > 0 {
			window[c]++
			// 当该字符在窗口中的数量满足 t 中的要求时，valid++
			if window[c] == need[c] {
				valid++
			}
		}

		// 尝试缩小窗口
		for valid == len(need) {
			// 更新最小子串
			if right-left < length {
				start = left
				length = right - left
			}

			// 移除左边的字符
			d := s[left]
			left++

			// 如果该字符在 t 中
			if need[d] > 0 {
				// 如果窗口中的该字符数量等于 t 中的数量，valid--
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	// 如果找不到符合条件的子串，返回 ""
	if length == math.MaxInt32 {
		return ""
	}

	// 返回最小覆盖子串
	return s[start : start+length]
}
