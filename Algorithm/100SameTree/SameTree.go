package algorithm

import (
	"ddkwing/arts/algorithm/structures"
)

type TreeNode = structures.TreeNode

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s != nil && t != nil {
		if s.Val == t.Val {
			return isSameTree(s.Left, t.Left) && isSameTree(t.Right, t.Right)
		}
		return false
	} else {
		return false
	}
}
