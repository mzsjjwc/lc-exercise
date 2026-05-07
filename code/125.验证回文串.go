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
	//思路是直接将其转换成两个数组,然后比较是否完全一致,需要提前处理字符串空格,大小写,还有非字母数字字符
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`) //定义正则
	s = reg.ReplaceAllString(s, "")            //替换
	//打印试试
	fmt.Printf("s = %s\n", s)
	chars := []rune(s)
	length := len(chars)
	start := 0
	end := length - 1
	for start < length && end >= 0 {
		if chars[start] != chars[end] {
			return false
		}
		start++
		end--
	}
	return true
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
