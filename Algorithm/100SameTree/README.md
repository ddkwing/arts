# 题目
Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

Example 1:


Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true

Example 2:


Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false

Example 3:


Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false

# 题目大意
这一题要求判断 2 颗树是否是完全相等的。

# 解题思路
利用递归思路来解，先判断跟节点，如果不同则两树不同，如果相同则递归两子树根节点，至少完成所有叶子遍历。