/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

package main

import "fmt"

// @lc code=start
func combine(n int, k int) [][]int {
	result := [][]int{}

	var backtrace func(list []int, begin int)
	backtrace = func(list []int, begin int) {
		if len(list) == k {
			l := make([]int, k)
			copy(l, list)
			result = append(result, l)
			return
		}

		for i := begin; i <= n; i++ {
			list = append(list, i)
			backtrace(list, i+1)
			list = list[:len(list)-1]
		}
	}
	backtrace([]int{}, 1)

	return result
}

// @lc code=end
func main() {
	fmt.Println(combine(4, 2))
}
