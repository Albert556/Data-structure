/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */
package golang

type ListNode struct {
	Val  int
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	a, b, c, l3 := 0, 0, 0, head
	for l1 != nil || l2 != nil || c != 0 {

		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		} else {
			a = 0
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		} else {
			b = 0
		}
		l3.Next = &ListNode{Val: (a + b + c) % 10}
		l3 = l3.Next
		c = (a + b + c) / 10
	}

	return head.Next
}

// @lc code=end
