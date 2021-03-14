/*
 * @lc app=leetcode.cn id=331 lang=golang
 *
 * [331] 验证二叉树的前序序列化
 */
package main

import "strings"

// @lc code=start
func isValidSerialization(preorder string) bool {
	s := strings.Split(preorder, ",")
	stack := []string{}
	for _, s := range s {
		stack = append(stack, s)
		for len(stack) >= 3 && stack[len(stack)-1] == "#" && stack[len(stack)-2] == "#" && stack[len(stack)-3] != "#" {
			stack = stack[:len(stack)-3]
			stack = append(stack, "#")
		}
	}

	return len(stack) == 1 && stack[0] == "#"
}

// @lc code=end
