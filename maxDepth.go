package main

import . "leetcode/tree"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	if leftMax > rightMax {
		return leftMax + 1
	} else {
		return rightMax + 1
	}
}