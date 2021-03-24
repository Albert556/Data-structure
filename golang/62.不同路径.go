/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */
package main

// @lc code=start
func uniquePaths(m int, n int) int {
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		result[i][0] = 1
	}
	for j := 0; j < n; j++ {
		result[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			result[i][j] = result[i-1][j] + result[i][j-1]
		}
	}

	return result[m-1][n-1]
}

// @lc code=end
