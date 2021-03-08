/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 */

package main

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
func levelOrder(root *TreeNode) [][]int {
	t := [][]int{}
	if root == nil {
		return t
	}
	quene := []*TreeNode{}
	quene = append(quene, root)
	for len(quene) > 0 {
		level := []int{}
		size := len(quene)
		for ; size > 0; size-- {
			if quene[0].Left != nil {
				quene = append(quene, quene[0].Left)
			}
			if quene[0].Right != nil {
				quene = append(quene, quene[0].Right)
			}
			level = append(level, quene[0].Val)
			quene = quene[1:]
		}
		t = append(t, level)
	}

	return t
}

// @lc code=end
