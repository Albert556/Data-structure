'''
Author: Albert
Date: 2021-01-14 20:59:59
LastEditors: Albert
LastEditTime: 2021-01-17 00:03:46
Desctiption: 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
'''


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def pathSum(self, root: TreeNode, sum: int) -> List[List[int]]:
        lists, path = [], []

        # 定义成类函数时需要更新属性，提交后多个实例之间不清除属性值
        def dfs(root, sum):
            if not root:
                return

            path.append(root.val)
            sum -= root.val

            if sum == 0 and not (root.left or root.right):
                # lists.append(path)将path加入lists中，不是复制，path更改，lists也会变
                lists.append(list(path))

            dfs(root.left, sum)
            dfs(root.right, sum)
            path.pop()

        dfs(root, sum)

        return lists


a, b, c, d, e, f, g, h, i, j = TreeNode(5), TreeNode(4), TreeNode(8), TreeNode(
    11), TreeNode(13), TreeNode(4), TreeNode(7), TreeNode(2), TreeNode(
        5), TreeNode(1)
a.left = b
a.right = c
b.left = d
b.right = None
c.left = e
c.right = f
d.left = g
d.right = h
e.left = None
e.right = None
f.left = i
f.right = j
g.left = None
g.right = None
h.left = None
h.right = None
i.left = None
i.right = None
j.left = None
j.right = None

sum = 22

print(Solution().pathSum(a, sum))
