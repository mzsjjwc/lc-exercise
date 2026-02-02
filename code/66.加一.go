/*
 * @lc app=leetcode.cn id=66 lang=golang
 * @lcpr version=30307
 *
 * [66] 加一
 */

package main

// @lc code=start
func plusOne(digits []int) []int {
	add := 0
	temp := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			temp = add + digits[i] + 1

		} else {
			temp = add + digits[i]

		}
		digits[i] = temp % 10
		add = temp / 10
	}
	if add == 1 {
		return append([]int{1}, digits...)
	} else {
		return digits
	}
	//更好的答案
	// for i := len(digits) - 1; i >= 0; i-- {
	// 	// 当前位加 1
	// 	digits[i]++
	// 	// 取模：如果是 10 则变 0，否则保持原样
	// 	digits[i] %= 10

	// 	// 如果当前位不是 0，说明没有进位了，直接返回
	// 	if digits[i] != 0 {
	// 		return digits
	// 	}
	// }

	// // 如果循环结束还没返回，说明是 [9, 9, 9] 这种情况
	// // 此时数组全变成了 [0, 0, 0]，需要在最前面补个 1
	// return append([]int{1}, digits...)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [4,3,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [9]\n
// @lcpr case=end

*/
