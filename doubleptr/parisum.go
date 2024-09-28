package doubleptr

import "sort"

//设计一个算法，找出数组中两数之和为指定值的所有整数对。一个数只能属于一个数对。
//1. 输入: nums = [5,6,5], target = 11
//输出: [[5,6]]
//2.输入: nums = [5,6,5,6], target = 11
//输出: [[5,6],[5,6]]

func pariSum(nums []int, target int) (ans [][]int) {
	sort.Ints(nums)

	//ans := make([][]int, 0)
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			ans = append(ans, []int{nums[left], nums[right]})
			left++
			right--
		}
	}
	return
}
