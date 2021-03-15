/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */
package main

import "fmt"

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// 用四个数维护上下左右四个边界，替代直接删减行列
	top, botton, left, right, result := 0, len(matrix), 0, len(matrix[0]), []int{}

	// 只要上下或者左右边界有重叠，证明全部遍历完成
	// 行列可以绑定为一组用来判断，行重叠时，遍历列的条件不成立，列同理
	for top < botton && left < right {
		for k := left; k < right; k++ {
			result = append(result, matrix[top][k])
		}
		top++
		for k := top; k < botton; k++ {
			result = append(result, matrix[k][right-1])
		}
		right--
		if top < botton && left < right {
			for k := right - 1; k >= left; k-- {
				result = append(result, matrix[botton-1][k])
			}
			botton--
			for k := botton - 1; k >= top; k-- {
				result = append(result, matrix[k][left])
			}
			left++
		}
	}

	return result
}

// @lc code=end

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
}
