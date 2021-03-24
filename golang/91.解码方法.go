/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */
package main

// @lc code=start
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := []int{}

	return dp[len(dp)-1]
}

// @lc code=end
