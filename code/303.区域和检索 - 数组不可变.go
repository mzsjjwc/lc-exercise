package main

/*
 * @lc app=leetcode.cn id=303 lang=golang
 * @lcpr version=30403
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	//初始化时提前计算好,数组长度取多一个是为了避免数组越界(-1)
	sums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sums[i+1] = nums[i] + sums[i]
	}
	return NumArray{sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

/*
// @lcpr case=start
// ["NumArray","sumRange","sumRange","sumRange"]\n[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]\n
// @lcpr case=end

*/

/**
给定一个整数数组  nums，处理以下类型的多个查询:
计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和 ，其中 left <= right
实现 NumArray 类：

NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int left, int right) 返回数组 nums 中索引 left 和 right 之间的元素的 总和 ，包含 left 和 right 两点（也就是 nums[left] + nums[left + 1] + ... + nums[right] )
*/
