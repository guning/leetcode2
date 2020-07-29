package main

import . "leetcode/tree"

func isSymmetric(root *TreeNode) bool {
	return root == nil || isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(l, r *TreeNode) bool {
	if r == nil || l == nil {
		return r == l
	}
	if r.Val != l.Val {
		return false
	}

	return isSymmetricHelper(l.Left, r.Right) && isSymmetricHelper(l.Right, r.Left)
}