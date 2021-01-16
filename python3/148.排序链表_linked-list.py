'''
Author: Albert
Date: 2021-01-14 15:15:31
LastEditors: Albert
LastEditTime: 2021-01-15 10:35:06
Desctiption: 排序链表，使用归并排序
'''

#
# @lc app=leetcode.cn id=148 lang=python3
#
# [148] 排序链表
#


# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def sortList(self, head: ListNode) -> ListNode:
        if not (head and head.next): return head

        # h 标示头结点，方便后续合并链表时便利
        # len 链表长度
        # v 表示当前合并的链表长度
        h, len, v = head, 0, 1
        while h:
            h, len = h.next, len + 1

        # 结果头指针，后续调整顺序可能导致头结点换位置
        res = ListNode()
        res.next = head

        # 开始归并排序
        # 当 v 和 len 一样时，代表归并排序完成，v代表本轮排序链表长度
        while v < len:

            # p作为排序节点的前驱节点，方便更改前驱节点的后继
            # h作为节点指针便利链表
            p, h = res, res.next
            # 本轮开始合并链表长度为 v 的链表
            while h:
                # 归并排序要两两合并
                # h1 是第一个要合并链表的头结点
                # i为本轮链表长度，i=0是找到链表的结尾
                h1, i = h, v
                while i and h:
                    h, i = h.next, i - 1
                # 如果 i 不为零结束遍历，证明 h = None，此时链表已经到头了，h1节点不够树，也没有h2链表了，本轮循环也可以结束了
                if i: break
                # 寻找h2头结点
                h2, i = h, v
                while i and h:
                    h, i = h.next, i - 1
                # 代表h1，h2的长度，v-i是因为可能存在节点尾部节点数量不够的情况，但是h2链表还是有节点需要归并
                c1, c2 = v, v - i
                # 开始归并，以其中一个链表归并完为结束
                while c1 and c2:
                    if h1.val <= h2.val:
                        p.next, h1, c1 = h1, h1.next, c1 - 1
                    else:
                        p.next, h2, c2 = h2, h2.next, c2 - 1
                    p = p.next
                # 其中一个链表归并完成，另一个链表还有节点没有完成归并，剩下的是有序的
                # 直接接到已合并的链表的结尾
                p.next = h1 if c1 else h2
                # 找到最后添加的链表的结尾
                while c1 > 0 or c2 > 0:
                    p, c1, c2 = p.next, c1 - 1, c2 - 1
                # 此时合并过的链表的为节点可能链接向某个节点，但是要让其只想后续未归并的链表才不会乱掉
                p.next = h
            # 一轮归并结束，下一轮归并链表的长度需要乘以2
            v *= 2

        return res.next


# @lc code=end
