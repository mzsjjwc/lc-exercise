/*
 * @lc app=leetcode.cn id=238 lang=golang
 * @lcpr version=30400
 *
 * [238] 除了自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
    //思路是利用前缀和,提前算出两个数组,每个数左侧的乘积,和每个数右侧的乘积
	//本题目中是要计算除该数字以外的所有数字乘积,所以可以拆分成这个数字左边的所有数字乘积乘以这个数字右边所有数字的乘积
	//比如说1,2,3,4
	//2这个位置的,答案应该是左边:1 右边3*4 总的1*3*4=12
	//那么知道怎么算了之后,开始顺序倒序各一遍即可
	answerL := make([]int, len(nums))
	answerL[0] = 1
	for i := 1; i < len(answerL); i++ {
		answerL[i] = answerL[i-1] * nums[i-1]
	}
	answerR := make([]int, len(nums))
	answerR[len(nums)-1] = 1
	for k := len(nums) - 2; k >= 0; k-- {
		answerR[k] = answerR[k+1] * nums[k+1]
	}
	//最后同位置互乘
	answer := make([]int, len(nums))
	for index := range nums {
		answer[index] = answerL[index] * answerR[index]
	}
	return answer
}
// @lc code=end



/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [-1,1,0,-3,3]\n
// @lcpr case=end

 */

