/*
 * @lc app=leetcode.cn id=219 lang=golang
 * @lcpr version=30401
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
    //这题初看,就是考查哈希表,
	//思路是:边遍历数组,边存,如果已经存在,则计算两者索引差值绝对值是否小于等于k,如果小于等于则直接返回,如果大于则更新索引然后继续找
	hasMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if val, ok := hasMap[v]; ok {
			if int(math.Abs(float64(val-i))) <= k {
				return true
			} else {
				hasMap[v] = i
			}
		} else {
			hasMap[v] = i
		}
	}
	return false
}
// @lc code=end



/*
// @lcpr case=start
// [1,2,3,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,0,1,1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1,2,3]\n2\n
// @lcpr case=end

 */

