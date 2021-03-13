/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */
package main

import "fmt"

// @lc code=start
// double stacks
// func calculate(s string) int {
// 	nums := []string{}
// 	ops := []string{}
// 	priority := map[string]int{
// 		"-": 1,
// 		"+": 1,
// 		"*": 2,
// 		"/": 2,
// 		"%": 2,
// 		"^": 3,
// 	}
// 	s = strings.Replace(s, " ", "", -1)
// 	s = strings.Replace(s, "(-", "(0-", -1)
// 	n := len(s)
// 	nums = append(nums, "0")
// 	for

// 	return 0
// }

// single stack
func calculate(s string) int {
	stack := []int{}
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	result := 0
	for _, v := range stack {
		result += v
	}

	return result
}

// @lc code=end

func main() {
	fmt.Println("h")
}
