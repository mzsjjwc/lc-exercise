/*
 * @lc app=leetcode.cn id=27 lang=golang
 * @lcpr version=30307
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	//因为如果正序遍历,如果删除,会影响后面的索引值
	//倒序遍历则不会
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,2,3]\n3\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,2,3,0,4,2]\n2\n
// @lcpr case=end

*/

