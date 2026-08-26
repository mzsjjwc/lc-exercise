/*
 * @lc app=leetcode.cn id=11 lang=golang
 * @lcpr version=30404
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
    //思路:利用双指针,每次移动短板,找到最大值
	p := 0
	q := len(height) - 1
	maxCapacity := 0
	for p < q {
		if height[p] > height[q] {
			tempArea := (q - p) * height[q]
			if tempArea > maxCapacity {
				maxCapacity = tempArea
			}
			q--
		} else {
			tempArea := (q - p) * height[p]
			if tempArea > maxCapacity {
				maxCapacity = tempArea
			}
			p++
		}
	}
	return maxCapacity
}
// @lc code=end



/*
// @lcpr case=start
// [1,8,6,2,5,4,8,3,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,1]\n
// @lcpr case=end

 */

