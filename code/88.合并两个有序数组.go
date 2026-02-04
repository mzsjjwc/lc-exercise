/*
 * @lc app=leetcode.cn id=88 lang=golang
 * @lcpr version=30307
 *
 * [88] 合并两个有序数组
 */

package main

import "fmt"

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	//定义三个指针,第一个位于第一个数组末尾,第二个位于第二个数组末尾,第三个则是位于最终数组的末尾
	p := m - 1
	q := n - 1
	end := len(nums1) - 1
	fmt.Println("初始 p = ", p, "q = ", q, "end = ", end)
	//只要第二个数组没有移动完,循环就继续
	for q >= 0 {
		fmt.Println("循环 p = ", p, "q = ", q)
		if p >= 0 && nums1[p] > nums2[q] {
			nums1[end] = nums1[p]
			p--
		} else {
			nums1[end] = nums2[q]
			q--
		}
		end--
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,0,0,0]\n3\n[2,5,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n[]\n0\n
// @lcpr case=end

// @lcpr case=start
// [0]\n0\n[1]\n1\n
// @lcpr case=end

*/
