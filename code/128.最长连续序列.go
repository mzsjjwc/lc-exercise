/*
 * @lc app=leetcode.cn id=128 lang=golang
 * @lcpr version=30401
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	//不能排序,只能用空间换时间
	//转换成map
	hasMap := make(map[int]bool)
	for _, v := range nums {
		hasMap[v] = true
	}

	//定义最长length
	maxLength := 0

	//开始遍历map,搜索最长序列
	for k := range hasMap {
		if !hasMap[k-1] { //这判断,只判断第一个,如果不是第一个就直接跳过了
			count := 0
			for hasMap[k] {
				count++
				k++
			}
			if count > maxLength {
				maxLength = count
			}
		}
	}

	return maxLength
}
// @lc code=end



/*
// @lcpr case=start
// [100,4,200,1,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,3,7,2,5,8,4,6,0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,1,2]\n
// @lcpr case=end

 */

