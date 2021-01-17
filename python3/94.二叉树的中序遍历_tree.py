'''
Author: Albert
Date: 2021-01-17 09:25:29
LastEditors: Albert
LastEditTime: 2021-01-17 09:30:23
Desctiption: 递归中序便利
'''

#
# @lc app=leetcode.cn id=94 lang=python3
#
# [94] 二叉树的中序遍历
#


# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        path = []

        def dfs_mid(root):
            if not (root): return

            dfs_mid(root.left)
            path.append(root.val)
            dfs_mid(root.right)

        dfs_mid(root)

        return path


# @lc code=end
