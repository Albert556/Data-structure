'''
Author: Albert
Date: 2021-01-14 10:31:41
LastEditors: Albert
LastEditTime: 2021-01-14 15:11:01
Desctiption:
'''

#
# @lc app=leetcode.cn id=143 lang=python3
#
# [143] 重排链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution:
    def reorderList(self, head: ListNode) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        if not (head and head.next and head.next.next): return head

        slow, fast = head, head
        while fast.next and fast.next.next:
            slow, fast = slow.next, fast.next.next

        # python slow.next = None是把下一个节点直接变成None而不是断开链接
        a, b = head, ListNode()
        b.next = slow.next
        slow.next = None

        if b.next and b.next.next:
            p, c, n = b, b.next, b.next.next
            while n:
                c.next = n.next
                n.next = p.next
                p.next = n
                n = c.next
            b = p.next
        else:
            b = b.next

        p, q, head = a.next, b.next, a
        while q:
            a.next = b
            b.next = p
            a, b = p, q
            p, q = p.next, q.next
        a.next = b
        b.next = p

        return head


# @lc code=end
# d = ListNode(4)
c = ListNode(3)
b = ListNode(2, c)
a = ListNode(1, b)
head = a

s = Solution()
s.reorderList(head)
while a:
    print(a.val)
    a = a.next
