/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */
package main

// @lc code=start
func setZeroes(matrix [][]int) {
	col0 := false

	for _, row := range matrix {
		if row[0] == 0 {
			col0 = true
		}
		for j := 1; j < len(matrix[0]); j++ {
			if row[j] == 0 {
				row[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 == true {
			matrix[i][0] = 0
		}
	}
}

// @lc code=end
