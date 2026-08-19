package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
 * @lc app=leetcode.cn id=125 lang=golang
 * @lcpr version=30403
 *
 * [125] 验证回文串
 */

//这里学到字符串就是rune字符组成的,可以直接将其转换成rune数组进行比较
//还有strings里面的ReplaceAll和ToLower函数
//还有正则表达式用法

// @lc code=start
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	p := 0
	q := len(s) - 1
	for p < q {
		for p < q && !isNumOrLetter(s[p]) {
			p++
		}
		for p < q && !isNumOrLetter(s[q]) {
			q--
		}
		if toLetter(s[p]) != toLetter(s[q]) {
			return false
		}
		p++
		q--
	}
	return true
}

func toLetter(t byte) byte {
	if t >= 'A' && t <= 'Z' {
		t = t + ('a' - 'A')
	}
	return t
}

func isNumOrLetter(t byte) bool {
	if t >= 'A' && t <= 'Z' || t >= 'a' && t <= 'z' || t >= '0' && t <= '9' {
		return true
	}
	return false
}


// @lc code=end

/*
// @lcpr case=start
// "A man, a plan, a canal: Panama"\n
// @lcpr case=end

// @lcpr case=start
// "race a car"\n
// @lcpr case=end

// @lcpr case=start
// " "\n
// @lcpr case=end

*/
