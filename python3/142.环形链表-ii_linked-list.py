'''
Author: Albert
Date: 2021-01-14 09:48:21
LastEditors: Albert
LastEditTime: 2021-01-14 10:27:52
Desctiption:
'''
#
# @lc app=leetcode.cn id=142 lang=python3
#
# [142] 环形链表 II
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def detectCycle(self, head: ListNode) -> ListNode:
        slow, fast = head, head
        while True:
            if fast and fast.next:
                slow, fast = slow.next, fast.next.next
            else:
                return None
            if slow == fast:
                break
        fast = head
        while slow != fast:
            slow, fast = slow.next, fast.next

        return fast
        # @lc code=end
