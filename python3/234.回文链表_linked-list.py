'''
Author: Albert
Date: 2021-01-15 15:01:31
LastEditors: Albert
LastEditTime: 2021-01-15 15:47:22
Desctiption:
'''

#
# @lc app=leetcode.cn id=234 lang=python3
#
# [234] 回文链表
#


# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        if not (head and head.next): return True

        slow, fast = head, head
        while fast.next and fast.next.next:
            slow, fast = slow.next, fast.next.next

        p, c, n = slow, slow.next, slow.next.next
        while n:
            c.next = n.next
            n.next = p.next
            p.next = n
            n = c.next
        slow, fast = head, p.next

        while fast:
            if slow.val != fast.val:
                return False
            slow, fast = slow.next, fast.next

        return True


# @lc code=end

b = ListNode(2)
a = ListNode(1)
a.next = b
Solution().isPalindrome(a)
