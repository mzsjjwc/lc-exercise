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
	//思路是,给每一个字符串进行排序,以排序后的字符串为键
	//值是字符串数组
	hasMap := make(map[string][]string)
	for _, v := range strs {
		temp := []byte(v)
		//排序
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		tempStr := string(temp)
		//如果存在,则添加,否则新建
		if val, ok := hasMap[tempStr]; ok {
			hasMap[tempStr] = append(val, v)
		} else {
			hasMap[tempStr] = []string{v}
		}
	}

	result := make([][]string, len(hasMap))

	j := 0
	for _, value := range hasMap {
		result[j] = value
		j++
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
