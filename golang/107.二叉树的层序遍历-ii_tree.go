/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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

func levelOrderBottom(root *TreeNode) [][]int {
	t := [][]int{}
	if root == nil {
		return t
	}
	quene := []*TreeNode{}
	quene = append(quene, root)

	for len(quene) > 0 {
		size := len(quene)
		level := []int{}
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
	length := len(t)
	for i := 0; i < length/2; i++ {
		t[i], t[length-i-1] = t[length-i-1], t[i]
	}

	return t
}

// @lc code=end

func main() {
	a, b, c, d, e := &TreeNode{Val: 3}, &TreeNode{Val: 9}, &TreeNode{Val: 20}, &TreeNode{Val: 15}, &TreeNode{Val: 7}
	a.Left = b
	a.Right = c
	c.Left = d
	c.Right = e
	fmt.Println(levelOrderBottom(a))
}
