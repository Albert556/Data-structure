<!--
 * @Author: Albert
 * @Date: 2021-01-17 00:53:40
 * @LastEditors: Albert
 * @LastEditTime: 2021-01-17 09:44:28
 * @Desctiption:
-->
# leetcode

## LinkList
## 解题技巧
1. 反转链表，用三指针
    ```python
    while n:
        c.next = n.next
        n.next = p.next
        p.next = n
        n = c.next
    ```
2. 链表排序<br>
   归并排序O(nlogn),O(1)，用迭代的方法，直接归并，省略了分治的过程

3. 环形链表
   * 链表是否有环用快慢指针，相遇有环
   * 二次相遇的位置是环的入口
## Tree
### 遍历
* 递归中序遍历
    ```python
    def dfs_mid(root):
        if not (root): return

        dfs_mid(root.left)
        # 逻辑代码
        path.append(root.val)
        dfs_mid(root.right)
    ```
