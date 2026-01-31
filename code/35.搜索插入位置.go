/*
 * @lc app=leetcode.cn id=35 lang=golang
 * @lcpr version=30307
 *
 * [35] 搜索插入位置
 */

package main

// @lc code=start
func searchInsert(nums []int, target int) int {
	//二分查找,顾名思义就是,每次取中间值,和目标值比较,
	//如果目标值大就往右边缩小范围,小就往左缩小一半范围
	//这里已经排好序所以不用提前排序
	//5个数的话 0 3 5 7 9 中间索引是 2
	//4个数的话 0 3 5 7 中间索引是1和2之间
	left := 0
	right := len(nums) - 1

	//先处理三种特殊情况,就是大于最大小于最小,等于最大
	if target <= nums[left] {
		return left
	}
	if target > nums[right] {
		return right + 1
	}
	if target == nums[right] {
		return right
	}

	//如果和中间值比,目标值大了,就移动左指针
	//中间值大了,就移动右指针,如果相等,则找到了
	//找到了则返回索引,如果没找到也就是比完了(左右指针相差1)都没有相等的,
	//则返回左指针加一
	for right-left > 1 {
		//右指针减去左指针除以二+左指针
		middle := (right-left)/2 + left
		if nums[middle] > target {
			right = middle
		} else if nums[middle] < target {
			left = middle
		} else {
			return middle
		}
	}

	//如果没找到,就只能是在中间了
	return left + 1

	// 标准答案
	// left := 0
	// right := len(nums) - 1
	// for left <= right {
	// 	mid := left + (right-left)/2
	// 	if nums[mid] == target {
	// 		return mid // 1. 直接命中
	// 	} else if nums[mid] < target {
	// 		left = mid + 1 // 2. 目标在右，左边界跨过 mid
	// 	} else {
	// 		right = mid - 1 // 3. 目标在左，右边界跨过 mid
	// 	}
	// }
	// return left // 4. 没找到，返回插入位置
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,5,6]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1,3,5,6]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,3,5,6]\n7\n
// @lcpr case=end

*/
