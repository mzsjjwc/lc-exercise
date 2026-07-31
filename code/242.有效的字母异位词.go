/*
 * @lc app=leetcode.cn id=242 lang=golang
 * @lcpr version=30307
 *
 * [242] 有效的字母异位词
 */

package main

import "sort"

// @lc code=start
func isAnagram(s string, t string) bool {
	c1 := statistic(s)
	c2 := statistic(t)
	for i:=0;i<26;i++{
		if c1[i] != c2[i]{
			return false
		}
	}
	return true
	
}

func statistic(str string)[]int{
	count := make([]int,26)
	for _,v:=range(str){
		index := v-'a'
		count[index]++
	}
	return count
}

// @lc code=end

/*
// @lcpr case=start
// "anagram"\n"nagaram"\n
// @lcpr case=end

// @lcpr case=start
// "rat"\n"car"\n
// @lcpr case=end

*/
