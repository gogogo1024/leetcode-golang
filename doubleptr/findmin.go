package doubleptr

//查找到有序递增数组旋转后数组中的最小值
// 输入: [3,4,5,1,2]
// 输出: 1

// 有序性：旋转前的数组是有序的，因此无论旋转后数组的形状如何，至少一半的数组仍然保持有序。
// 分割策略：通过比较 nums[mid] 和 nums[right]，可以判断当前数组的哪一侧包含最小值：
//	1. 如果 nums[mid] > nums[right]，最小值一定在右半部分，因为左半部分是有序的。
//	2. 如果 nums[mid] <= nums[right]，则最小值在左半部分或可能是 mid。
//缩小范围：每次迭代都将搜索范围缩小一半，保证了算法的高效性，时间复杂度为 O(log n)。

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
