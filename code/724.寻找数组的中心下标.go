/*
 * @lc app=leetcode.cn id=724 lang=golang
 * @lcpr version=30403
 *
 * [724] 寻找数组的中心下标
 */

// @lc code=start
func pivotIndex(nums []int) int {
	fmt.Println(nums)
	//定义两个前缀和数组,分别从两侧计算前缀和
	leftPreSum := make([]int, len(nums))
	rightPreSum := make([]int, len(nums))
	lSum := 0
	rSum := 0
	for i, v := range nums {
		lSum += v
		leftPreSum[i] = lSum
		// fmt.Printf("lSum = %d,leftPreSum[%d] = %d\n", lSum, i, leftPreSum[i])
	}
	fmt.Println()
	for j := len(nums) - 1; j >= 0; j-- {
		rSum += nums[j]
		rightPreSum[j] = rSum
		// fmt.Printf("rSum = %d,rightPreSum[%d] = %d\n", rSum, j, rightPreSum[j])
	}

	//循环然后找到中心点
	for k := 0; k < len(nums); k++ {
		var left, right int //定义左边和和右边和
		//这里通过两个边界处理让循环到达了最左边和最右边
		if k > 0 {
			//如果k大于零,则减一,这里判断避免越界
			left = leftPreSum[k-1]
		}
		if k < len(nums)-1 {
			//这里判断同样是为了避免越界
			right = rightPreSum[k+1]
		}

		if left == right {
			return k
		}
	}
	return -1

}

// @lc code=end

/*
// @lcpr case=start
// [1,7,3,6,5,6]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,-1]\n
// @lcpr case=end

*/
