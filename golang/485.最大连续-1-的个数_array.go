/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续 1 的个数
 */

package main

import "fmt"

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	max, count := 0, 0
	for _, n := range nums {
		if n != 0 {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}
	if count != 0 && count > max {
		max = count
		count = 0
	}

	return max
}

// @lc code=end
func main() {
	nums := []int{1, 0, 1, 1, 0, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
