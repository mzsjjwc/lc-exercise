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
	chars1 := []rune(s)
	chars2 := []rune(t)
	if len(chars1) != len(chars2) {
		return false
	}
	sort.Slice(chars1, func(i, j int) bool {
		return chars1[i] < chars1[j]
	})
	sort.Slice(chars2, func(i, j int) bool {
		return chars2[i] < chars2[j]
	})
	length := len(chars1)
	for a := 0; a < length; a++ {
		if chars1[a] != chars2[a] {
			return false
		}
	}
	return true
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
