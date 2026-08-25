/*
 * @lc app=leetcode.cn id=15 lang=golang
 * @lcpr version=30403
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
    //思路是:先排序,固定一个数,然后当成双指针做,注意去重,去重的思路就是碰到相同元素直接跳过
	slices.Sort(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		//利用短路特性防止越界
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		p := i + 1
		q := len(nums) - 1
		for p < q {
			if nums[p]+nums[q]+nums[i] == 0 {
				result = append(result, []int{nums[p], nums[q], nums[i]})
				for p < q && nums[p] == nums[p+1] {
					p++
				}
				for p < q && nums[q] == nums[q-1] {
					q--
				}
				p++
				q--
				continue
			} else if nums[p]+nums[q]+nums[i] < 0 {
				p++
			} else if nums[p]+nums[q]+nums[i] > 0 {
				q--
			}
		}
	}
	return result
}
// @lc code=end



/*
// @lcpr case=start
// [-1,0,1,2,-1,-4]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n
// @lcpr case=end

 */

