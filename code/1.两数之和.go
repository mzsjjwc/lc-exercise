/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=30307
 *
 * [1] 两数之和
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	//先将数组转化成map
	hasMap := make(map[int]int, len(nums))
	//然后直接找taget－当前值是否已经在map中了，如果找到了则直接返回下标
	for i, v := range nums {
		temp := target - v
		if val, ok := hasMap[temp]; ok {
			return []int{i, val}
		} else {
			hasMap[v] = i
		}
	}
	return []int{}
}

// @lc code=end

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/
