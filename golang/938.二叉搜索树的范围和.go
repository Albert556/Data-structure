/*
 * @lc app=leetcode.cn id=938 lang=golang
 *
 * [938] 二叉搜索树的范围和
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// dfs
// func rangeSumBST(root *TreeNode, low int, high int) int {
// 	sum := 0
// 	stack := []*TreeNode{}
// 	stack = append(stack, root)

// 	for len(stack) > 0 {
// 		s := stack[len(stack)-1]
// 		stack = stack[0 : len(stack)-1]
// 		if s != nil {
// 			if low <= s.Val && s.Val <= high {
// 				sum += s.Val
// 			}
// 			if s.Val < high {
// 				stack = append(stack, s.Right)
// 			}
// 			if low < s.Val {
// 				stack = append(stack, s.Left)
// 			}
// 		}
// 	}

// 	return sum
// }

//
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	if high < root.Val {
		return rangeSumBST(root.Left, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

// @lc code=end

func main() {
	a := []interface{}{10, 5, 15, 3, 7, nil, 18}
	low := 7
	high := 15
	aa := []*TreeNode{}
	for _, v := range a {
		if v == nil {
			aa = append(aa, nil)
		} else {
			node := &TreeNode{Val: (v).(int), Left: nil, Right: nil}
			aa = append(aa, node)
		}
	}
	aa[0].Left = aa[1]
	aa[0].Right = aa[2]
	aa[1].Left = aa[3]
	aa[1].Right = aa[4]
	aa[2].Left = aa[5]
	aa[2].Right = aa[6]
	fmt.Println(rangeSumBST(aa[0], low, high))
}
