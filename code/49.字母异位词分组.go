/*
 * @lc app=leetcode.cn id=49 lang=golang
 * @lcpr version=30307
 *
 * [49] 字母异位词分组
 */
package main

import "sort"

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	//思路：利用排序后的字符串作为键，将异位词收集到同一个key下
	hasMap := make(map[string][]string)
	for _, v := range strs {
		r := []rune(v) 
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		newStr := string(r)
		hasMap[newStr] = append(hasMap[newStr], v)

	}
	//二维数组，每一个元素里面有一个[]string,
	result := make([][]string,0,len(hasMap))
	for _,vv := range(hasMap){
		result = append(result,vv)

	}
	return result

}

// @lc code=end

/*
// @lcpr case=start
// ["eat","tea","tan","ate","nat","bat"]\n
// @lcpr case=end

// @lcpr case=start
// [""]\n
// @lcpr case=end

// @lcpr case=start
// ["a"]\n
// @lcpr case=end

*/

//标准写法有待研究
// func groupAnagrams(strs []string) [][]string {
//     // Key 是长度为 26 的数组，记录 a-z 出现的次数
//     hasMap := make(map[[26]int][]string)

//     for _, v := range strs {
//         cnt := [26]int{}
//         for i := 0; i < len(v); i++ {
//             cnt[v[i]-'a']++
//         }
//         // 直接以数组作为 Key，将原字符串加入对应的切片
//         hasMap[cnt] = append(hasMap[cnt], v)
//     }

//     // 提取结果
//     result := make([][]string, 0, len(hasMap))
//     for _, value := range hasMap {
//         result = append(result, value)
//     }
//     return result
// }
