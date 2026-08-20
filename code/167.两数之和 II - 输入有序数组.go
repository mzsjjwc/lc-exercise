package main

/*
 * @lc app=leetcode.cn id=167 lang=golang
 * @lcpr version=30403
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
		//可以用哈希表做,也可以用双指针,双指针空间复杂度更低,时间复杂度和哈希表一样都是On,空间复杂度则是O1,思路是,因为是递增数组,所以往左就是减小,往右就是增大
	p := 0
	q := len(numbers) - 1
	for p < q {
		if numbers[p]+numbers[q] > target {
			q--
		} else if numbers[p]+numbers[q] < target {
			p++
		} else if numbers[p]+numbers[q] == target {
			return []int{p + 1, q + 1}
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
// [2,3,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [-1,0]\n-1\n
// @lcpr case=end

*/
