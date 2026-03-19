/*
 * @lc app=leetcode.cn id=347 lang=golang
 * @lcpr version=30400
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
    	//先按频率转换成map
	hasMap := make(map[int]int)
	for _, v := range nums {
		hasMap[v]++
	}

	//再用桶排序,注意这里桶是二维数组,一维数组无法解决相同频率的问题(频率多少就放哪个桶里,比如出现了2次,则放在bucket[2]里,最后倒序遍历即可)
	buckets := make([][]int, len(nums)+1)
	for k, v := range hasMap {
		buckets[v] = append(buckets[v], k)
	}

	result := make([]int, 0)

	//最后再倒序取出k个
	for i := len(buckets) - 1; i >= 0; i-- {
		if len(buckets[i]) == 0 {
			continue
		}
		for _, num := range buckets[i] {
			result = append(result, num)
			//确保只取k个
			if len(result) == k {
				return result
			}
		}
	}

	return result
}
// @lc code=end



/*
// @lcpr case=start
// [1,1,1,2,2,3]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1,2,1,2,3,1,3,2]\n2\n
// @lcpr case=end

 */

