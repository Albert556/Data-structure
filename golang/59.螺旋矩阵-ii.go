/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */
package main

import "fmt"

// @lc code=start
func generateMatrix(n int) [][]int {
	nn := n * n
	resullt := make([][]int, n)
	for i := 0; i < n; i++ {
		resullt[i] = make([]int, n, n)
	}
	top, botton, left, right := 0, n, 0, n
	k := 1
	for {
		if k <= nn {
			for i := left; i < right; i++ {
				resullt[top][i] = k
				k++
			}
			top++
		} else {
			break
		}
		if k <= nn {
			for i := top; i < botton; i++ {
				resullt[i][right-1] = k
				k++
			}
			right--
		} else {
			break
		}
		if k <= nn {
			for i := right - 1; i >= left; i-- {
				resullt[botton-1][i] = k
				k++
			}
			botton--
		} else {
			break
		}
		if k <= nn {
			for i := botton - 1; i >= top; i-- {
				resullt[i][left] = k
				k++
			}
			left++
		} else {
			break
		}
	}

	return resullt
}

// @lc code=end

func main() {
	fmt.Println(generateMatrix(3))
}
