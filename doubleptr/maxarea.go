package doubleptr

// 给定一个长度为 n 的整数数组 height 。
// 有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i])。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。
//输入: nums = [1,8,6,2,5,4,8,3,7]
//输出: 49

// 设两指针 i , j ，指向的水槽板高度分别为 h[i], h[j] ，此状态下水槽面积为 S(i, j)。
// 由于可容纳水的高度由两板中的 短板决定，因此可得如下面积公式：S(i, j) = min(h[i], h[j])×(j−i)
// 假设 h[j] > h[i]
//在每个状态下，无论长板或短板向中间收窄一格，都会导致水槽底边宽度−1 变短：
// 1. 若向内移动短板 i，水槽的短板 min(h[i], h[j]) 可能变大，因此下个水槽的面积可能增大。
// 2. 若向内移动长板 j，水槽的短板 min(h[i], h[j]) 不变或变小，因此下个水槽的面积一定变小
// 因此每次向内移动短板，才会让下个水槽面积可能变大

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	for i < j {
		if height[i] < height[j] {
			res = max(res, (j-i)*height[i])
			i++
		} else {
			res = max(res, (j-i)*height[j])
			j--
		}
	}
	return res
}
