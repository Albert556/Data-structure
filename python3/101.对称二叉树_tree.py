'''
Author: Albert
Date: 2021-01-16 09:51:43
LastEditors: Albert
LastEditTime: 2021-01-16 19:30:56
Desctiption:
'''
#
# @lc app=leetcode.cn id=101 lang=python3
#
# [101] 对称二叉树
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:

# @lc code=end

