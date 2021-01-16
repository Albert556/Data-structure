'''
Author: Albert
Date: 2021-01-15 10:35:32
LastEditors: Albert
LastEditTime: 2021-01-16 09:44:24
Desctiption:
'''

#
# @lc app=leetcode.cn id=206 lang=python3
#
# [206] 反转链表
#


# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    # 迭代
    # def reverseList(self, head: ListNode) -> ListNode:
    #     if not (head and head.next): return head

    #     p, c, n = ListNode(0), head, head.next
    #     p.next = head

    #     while n:
    #         c.next = n.next
    #         n.next = p.next
    #         p.next = n
    #         n = c.next

    #     return p.next

    # 递归
    def reverseList(self, head: ListNode) -> ListNode:
        if not (head and head.next): return head

        p = ListNode(0)
        p.next = None
        p = self.reverseList(head.next)
        head.next.next = head
        head.next = None

        return p


# @lc code=end
b = ListNode(2)
a = ListNode(1)
a.next = b
Solution().reverseList(a)
