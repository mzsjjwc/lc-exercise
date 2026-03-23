/*
 * @lc app=leetcode.cn id=238 lang=golang
 * @lcpr version=30400
 *
 * [238] 除了自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
    left := make([]int, len(nums))
	left[0] = 1 //由于索引0的元素左边没有元素,乘积为1
	//先计算出左边每个数的左边乘积(不包含自身)
	for i := 1; i < len(nums); i++ {
		//利用前缀积的思想,不重复计算,这里要这样看,每一次遍历都等于上一次的乘积(left[i-1])乘以这一次的乘数(nums[i-1])
		left[i] = left[i-1] * nums[i-1]
	}
	//再和右边每个元素的乘积相乘就能得到每个位置除了自身的乘积,这样只需要算三个循环,时间复杂度为3n,为O(n)
	//这里需要从右往左,由于最后一个位置的右边乘积为1,所以最后一个位置不变即可
	right := make([]int, len(nums))
	right[len(nums)-1] = 1
	for j := len(nums) - 2; j >= 0; j-- {
		right[j] = right[j+1] * nums[j+1]
	}
	result := make([]int, len(nums))
	for k := 0; k < len(nums); k++ {
		result[k] = left[k] * right[k]
	}
	return result
}
// @lc code=end



/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [-1,1,0,-3,3]\n
// @lcpr case=end

 */

