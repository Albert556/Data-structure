/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var f func(root *TreeNode, sum int) bool
	f = func(root *TreeNode, sum int) bool {
		if root == nil {
			return false
		}
		if root.Left == nil && root.Right == nil {
			return sum == root.Val
		}

		return f(root.Left, sum-root.Val) || f(root.Right, sum-root.Val)
	}

	return f(root, targetSum)
}

// @lc code=end
func main() {
	a, b, c, d, e, f, g, h, i := &TreeNode{Val: 5}, &TreeNode{Val: 4}, &TreeNode{Val: 8}, &TreeNode{Val: 11}, &TreeNode{Val: 13}, &TreeNode{Val: 4}, &TreeNode{Val: 7}, &TreeNode{Val: 2}, &TreeNode{Val: 1}
	a.Left = b
	a.Right = c
	b.Left = d
	c.Left = e
	c.Right = f
	d.Left = g
	d.Right = h
	f.Right = i
	fmt.Println(hasPathSum(a, 22))
}
