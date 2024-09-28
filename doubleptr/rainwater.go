package doubleptr

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子， 下雨之后能接多少雨水。
//输入: height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6
//Trapping Rain Water

func trapRainWater(height []int) (water int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			water += leftMax - height[left]
			left++
		} else {
			water += rightMax - height[right]
			right--
		}
	}
	return
}
