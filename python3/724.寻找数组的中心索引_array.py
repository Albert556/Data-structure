'''
Author: Albert
Date: 2021-01-19 22:16:08
LastEditors: Albert
LastEditTime: 2021-01-19 22:33:07
Desctiption:
'''

#
# @lc app=leetcode.cn id=724 lang=python3
#
# [724] 寻找数组的中心索引
#


# @lc code=start
class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        s = sum(nums)
        leftsum = 0
        for i, x in enumerate(nums):
            if leftsum == s - x - leftsum: return i
            leftsum += x
        return -1


# @lc code=end
