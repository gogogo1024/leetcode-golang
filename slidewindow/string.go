package slidewindow

//给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。
//子串 (Substring) 是指字符串中连续的字符组成的片段。
//例如, 对于字符串"abcxyz", 它的子串有"bcx"、"xyz" 等.
// 子串必须是连续的, 子序列可以是不连续的。
// 子串是字符串的一部分, 子序列不必是字符串的一部分。
// 任何子串都可以看成是一个子序列, 但是子序列不一定是子串。
//1
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是"abc"，所以其长度为 3。
//2
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是"b"，所以其长度为 1。
//3 输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。

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
