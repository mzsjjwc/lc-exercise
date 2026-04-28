/*
 * @lc app=leetcode.cn id=560 lang=golang
 * @lcpr version=30403
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	//TODO:本题写不出来,只好抄写,然后注释了
    count := 0 //用来统计子数组个数
    preSum := 0 //前缀和
    m := make(map[int]int) //k存储前缀和,v存储前缀和出现次数
    m[0] = 1 //这里必须特殊处理,如果不处理会漏掉所有等于k的次数
    for _, num := range nums {
        preSum += num //计算每一步的前缀和
        if val, ok := m[preSum-k]; ok { //这个实际上是在找,两个前缀和差值等于k的另外一个前缀和,因为preSum是当前前缀和,这个关键的要理解两个前缀和之间的差值的算法{sum}(j, i) = pre[i] - pre[j-1]
            count += val //如果找到了,则加上该次数,由于每次找的都是不同的另外一个前缀和,所以这个不存在重复加的情况
        }
        m[preSum]++ //这里前缀和出现次数正常累加即可
    }
    //返回最终统计结果
    return count
}
// @lc code=end



/*
// @lcpr case=start
// [1,1,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n3\n
// @lcpr case=end

 */

