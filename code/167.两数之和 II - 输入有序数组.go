package main

/*
 * @lc app=leetcode.cn id=167 lang=golang
 * @lcpr version=30403
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	//思路:这里利用有序数组来用双指针,因为是非递减数组,所以往左是变小,往右是变大
	//也可以用哈希表,但是时间复杂度和空间复杂度都比这个高,现在这个是O(1)
	left := 0
	right := len(numbers) - 1
	for left < right { //这里条件就是,所有的数都比对完了,然后这里不能用相等,因为数组个数不确定,如果是单数则可以相等,如果是双数则永远无法相等
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
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
