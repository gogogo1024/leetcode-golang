package doubleptr

import (
	"math"
)

//给定一个整数数组，编写一个函数，找出索引 m 和 n，只要将索引区间 [m, n] 的元素
//排好序，整个数组就是有序的。注意：n−m 尽量最小，也就是说，找出符合条件的最
//短序列。函数返回值为 [m, n]，若不存在这样的 m 和 n（例如整个数组是有序的），请
//返回 [−1,−1]

// 输入: [1,2,4,7,10,11,7,12,6,7,16,18,19]
// 输出: [3,9]
func findUnsortedSubarray(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}

	// Step 1: 从左到右找到第一个逆序对
	m := -1
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			m = i - 1
			break
		}
	}
	if m == -1 { // 数组已经是有序的
		return []int{-1, -1}
	}

	// Step 2: 从右到左找到第一个逆序对
	nn := -1
	for i := n - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			nn = i + 1
			break
		}
	}

	// Step 3: 找到 nums[m:n+1] 区间内的最小值和最大值
	subMin := math.MaxInt
	subMax := math.MinInt
	for i := m; i <= nn; i++ {
		if nums[i] < subMin {
			subMin = nums[i]
		}
		if nums[i] > subMax {
			subMax = nums[i]
		}
	}

	// Step 4: 扩展 m 和 n 的边界
	for m > 0 && nums[m-1] > subMin {
		m--
	}
	for nn < n-1 && nums[nn+1] < subMax {
		nn++
	}

	return []int{m, nn}
}
