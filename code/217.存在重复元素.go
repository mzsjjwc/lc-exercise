/*
 * @lc app=leetcode.cn id=217 lang=golang
 * @lcpr version=30307
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	numMap := make(map[int]interface{}, len(nums)) //预分配内存节省开销
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			return true
		} else {
			numMap[v] = struct{}{}
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,3,3,4,3,2,4,2]\n
// @lcpr case=end

*/
