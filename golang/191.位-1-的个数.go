/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */
package main

// @lc code=start
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num%2 == 1 {
			count++
		}
		num /= 2
	}

	return count
}

// @lc code=end
