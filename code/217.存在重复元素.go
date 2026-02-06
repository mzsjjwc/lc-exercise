/*
 * @lc app=leetcode.cn id=217 lang=golang
 * @lcpr version=30307
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	hasMap := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := hasMap[v]; ok {
			return true //如果有重复的,就提前跳出
		} else {
			hasMap[v] = struct{}{} //这里是空结构体字面量可以用[]int{}来类比,[]int可以替换成struct{},这样理解就可以了
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

