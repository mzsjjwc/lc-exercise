/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=30307
 *
 * [1] 两数之和
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	//创建hashmap
	hasMap := make(map[int]int, len(nums))
	for i, v := range nums {
		//看看map里有没有目标值，如果有就返回没有就加进去当前的值和索引
		if val, ok := hasMap[target-v]; ok {
			return []int{val, i}
		} else {
			hasMap[v] = i
		}
	}
	return nil
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
