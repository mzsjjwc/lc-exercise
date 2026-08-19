/*
 * @lc app=leetcode.cn id=128 lang=golang
 * @lcpr version=30401
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	//思路：因为是未排序，并且要求时间复杂度O(n)，那就不能直接排序了，只能用其他方法
	//可以转换成哈希表，然后找第一个，跳过所有非第一个元素的
	hasMap := make(map[int]bool)
	fmt.Println(hasMap)
	for _, v := range nums {
		hasMap[v] = true
	}
	fmt.Println(hasMap)
	maxLength := 0
	for key := range hasMap {
		if !hasMap[key-1] {
			tempKey := key
			count := 0
			for hasMap[tempKey] {
				count++
				tempKey++
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

