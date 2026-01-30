/*
 * @lc app=leetcode.cn id=26 lang=golang
 * @lcpr version=30307
 *
 * [26] 删除有序数组中的重复项
 */
package main

import "slices"

// @lc code=start
func removeDuplicates(nums []int) int {
	//先排序
	slices.Sort(nums)
	slow := 0
	for quick := slow; quick < len(nums); quick++ {
		//1,2,3,4,4,4,5
		//1,2,3,4,5,4,5
		//1,2,3,4,5
		//q =
		//s =
		if nums[quick] != nums[slow] {
			//如果快指针指向的元素不等于慢指针指向的元素,那么说明是新元素
			//slow则永远代表最新的一个元素
			slow++
			nums[slow] = nums[quick]
		}
	}
	//去除掉slow后面的元素,左闭右开,如果想保留nums[slow]则需要slow+1
	// nums = append(nums[:slow+1])
	return slow + 1
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,1,2,2,3,3,4]\n
// @lcpr case=end

*/
