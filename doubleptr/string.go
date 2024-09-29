package doubleptr

// 和slidewindow/string 一样但是解法是使用双指针
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

func lengthOfLongestSubstring(s string) int {
	// 用于记录每个字符的上一次出现位置，数组大小为 256 表示 ASCII 字符集
	pos := [256]int{} // 初始值为 0，表示每个字符未出现时的默认位置
	n := len(s)
	ans, left := 0, 0 // ans 是最长子串的长度，left 是滑动窗口的左边界

	// 右指针从 0 开始遍历整个字符串
	for right := 0; right < n; right++ {
		ch := s[right] // 获取当前字符
		// 更新左边界 left，确保它在重复字符的上一次出现位置之后
		if pos[ch] > left {
			left = pos[ch] // 更新 left 指针到上次重复字符的位置的右边
		}
		// 记录字符的最新出现位置，+1 是为了区分字符是否从未出现过
		pos[ch] = right + 1
		// 更新无重复字符子串的最大长度
		ans = max(ans, right-left+1)
	}

	return ans
}
