/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */
package main

import "fmt"

type ListNode struct {
	val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := &ListNode{Next: head}
	pre := p
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	cur, next := pre.Next, pre.Next.Next
	for i := left; i < right; i++ {
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
		next = cur.Next
	}

	return p.Next
}

// @lc code=end
func main() {
	a, left, right := []int{1, 2, 3, 4, 5}, 2, 4
	b := &ListNode{}
	root := b
	for _, v := range a {
		c := &ListNode{val: v}
		b.Next = c
		b = b.Next
	}
	fmt.Println(reverseBetween(root.Next, left, right))
}
