package main

import (
	"fmt"
	. "leetcode/tree"
)

func findTarget(root *TreeNode, k int) bool {
	return fn(root, root, k)

}

func fn (tn1, tn2 *TreeNode, k int) bool {
	if tn1 == nil || tn2 == nil {
		return false
	}

	if tn1.Val + tn2.Val == k {
		if tn1 == tn2 {
			return fn(tn1.Left, tn1.Right, k)
		}
		return true
	} else if tn1.Val + tn2.Val < k {
		if tn1 == tn2 {
			return fn(tn1, tn1.Right, k)
		}
		return fn(tn1, tn2.Right, k) || fn(tn1.Right, tn2, k)
	} else {
		if tn1 == tn2 {
			return fn(tn1, tn1.Left, k)
		}
		return fn(tn1, tn2.Left, k) || fn(tn1.Left, tn2, k)
	}
	return false
}


func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	T := GenerateRadBlackTree(nums)
	fmt.Println(findTarget(T.Root, 9))
}
