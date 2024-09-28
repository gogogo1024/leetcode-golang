package doubleptr

//给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使
//得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。
//必须在不使用库内置的 sort 函数的情况下解决这个问题。
//输入: nums = [2,0,2,1,1,0]
//输出: [0,0,1,1,2,2]

// 如果当前元素是 0，交换 nums[i] 和 nums[low]，将 0 移动到前面
// 如果当前元素是 2，交换 nums[i] 和 nums[high]，将 2 移动到后面
// 如果当前元素是 1， 直接跳过 1，因为它已经在正确的位置

// sortColors 对三色数组进行原地排序
func sortColor(nums []int) []int {
	low, high := 0, len(nums)-1
	i := 0
	for i <= high {
		switch nums[i] {
		case 0:
			nums[i], nums[low] = nums[low], nums[i]
			low++
			i++
		case 2: // [10] 如果当前元素是 2
			nums[i], nums[high] = nums[high], nums[i]
			high--
		default:
			i++
		}
	}
	return nums
}
