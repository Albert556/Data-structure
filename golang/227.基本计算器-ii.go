/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */
package main

import (
	"fmt"
	"math"
	"strings"
)

// @lc code=start
// double stacks
func calculate(s string) int {
	nums := []int{}
	ops := []byte{}
	priority := map[byte]int{
		'-': 1,
		'+': 1,
		'*': 2,
		'/': 2,
		'%': 2,
		'^': 3,
	}
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "(-", "(0-", -1)
	var calc func()
	calc = func() {
		if len(nums) < 2 {
			return
		}
		if (len(ops)) == 0 {
			return
		}
		b := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		a := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		ans := 0
		if op == '+' {
			ans = a + b
		} else if op == '-' {
			ans = a - b
		} else if op == '*' {
			ans = a * b
		} else if op == '/' {
			ans = a / b
		} else if op == '^' {
			ans = int(math.Pow(float64(a), float64(b)))
		} else if op == '%' {
			ans = a % b
		}
		nums = append(nums, ans)

	}

	nums = append(nums, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			ops = append(ops, s[i])
		} else if s[i] == ')' {
			for len(ops) > 0 {
				if ops[len(ops)-1] != '(' {
					calc()
				} else {
					ops = ops[:len(ops)-1]
					break
				}
			}
		} else {
			if '0' <= s[i] && s[i] <= '9' {
				num := int(s[i] - '0')
				j := i + 1
				for ; j < len(s) && '0' <= s[j] && s[j] <= '9'; j++ {
					num = num*10 + int(s[j]-'0')
				}
				nums = append(nums, num)
				i = j - 1
			} else {
				if len(ops) > 0 && ops[len(ops)-1] != '(' {
					preOps := ops[len(ops)-1]
					if priority[preOps] >= priority[s[i]] {
						calc()
					}
				}
				ops = append(ops, s[i])
			}
		}
	}

	for len(ops) > 0 {
		calc()
	}
	return nums[len(nums)-1]
}

func calc(nums []int, ops []byte) {
	if len(nums) < 2 {
		return
	}
	if (len(ops)) == 0 {
		return
	}
	b := nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	a := nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	op := ops[len(ops)-1]
	ops = ops[:len(ops)-1]
	ans := 0
	if op == '+' {
		ans = a + b
	} else if op == '-' {
		ans = a - b
	} else if op == '*' {
		ans = a * b
	} else if op == '/' {
		ans = a / b
	} else if op == '^' {
		ans = int(math.Pow(float64(a), float64(b)))
	} else if op == '%' {
		ans = a % b
	}
	nums = append(nums, ans)

}

// single stack
// func calculate(s string) int {
// 	stack := []int{}
// 	preSign := '+'
// 	num := 0
// 	for i, ch := range s {
// 		isDigit := '0' <= ch && ch <= '9'
// 		if isDigit {
// 			num = num*10 + int(ch-'0')
// 		}
// 		if !isDigit && ch != ' ' || i == len(s)-1 {
// 			switch preSign {
// 			case '+':
// 				stack = append(stack, num)
// 			case '-':
// 				stack = append(stack, -num)
// 			case '*':
// 				stack[len(stack)-1] *= num
// 			default:
// 				stack[len(stack)-1] /= num
// 			}
// 			preSign = ch
// 			num = 0
// 		}
// 	}
// 	result := 0
// 	for _, v := range stack {
// 		result += v
// 	}

// 	return result
// }

// @lc code=end

func main() {
	fmt.Println(calculate("-1+ (31+2)*3"))
}
