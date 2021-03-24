/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */
package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// @lc code=start
func evalRPN(tokens []string) int {
	nums := []int{}
	re := "[0-9]+"
	for _, token := range tokens {
		if t, _ := regexp.MatchString(re, token); t {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		} else {
			b := nums[len(nums)-1]
			a := nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			switch token {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			}
		}
	}
	return nums[0]
}

// @lc code=end
func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}
