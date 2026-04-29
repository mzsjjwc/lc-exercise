/*
 * @lc app=leetcode.cn id=974 lang=golang
 * @lcpr version=30403
 *
 * [974] 和可被 K 整除的子数组
 */

// @lc code=start
func subarraysDivByK(nums []int, k int) int {
    //理解题目:这题又是前缀和题目
	//元素之和可以被k整除也就是x%k == 0的
	//肯定还是利用前缀和公式{sum}(j, i) = pre[i] - pre[j-1]
	//(pre[i] - pre[j-1]) % k == 0
	//根据同余定理,PrefixSum[j] \pmod K == PrefixSum[i-1] \pmod K
	//也就是说只要这两个前缀和的余数相同,则成立
	//关键思路:需要找到(pre[i] - pre[j-1]) % k == 0的两个前缀和
	//然后需要把公式变体为pre[i]%k == pre[j-1]%k,也就是两者余数相等
	//找两者余数相等的前缀和,遍历,然后找余数相等的前缀和出现次数,加起来,就可以了,详细可以看724题目
	preSum := 0
	count := 0
	m := make(map[int]int)
	m[0] = 1
	for _,v := range(nums){
		preSum += v
		y := (preSum % k + k) % k //这个地方本来取模就行了,但是为了保证负数取模正常,所以加k再对k取模,结果就能正常,或者另外一种方法,加个判断如果是负数就加个k补正
		if val,ok:=m[y];ok{
			count += val
		}
		m[y]++
		
	}
	return count
}
// @lc code=end



/*
// @lcpr case=start
// [4,5,0,-2,-3,1]\n5\n
// @lcpr case=end

// @lcpr case=start
// [5]\n9\n
// @lcpr case=end

 */

