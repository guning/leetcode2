package main

import (
	. "leetcode/list"
	"sort"
)

func mergeKLists(lists []*ListNode) *ListNode {
	var resArr []int
	for _, v := range lists {
		for v != nil {
			resArr = append(resArr, v.Val)
			v = v.Next
		}
	}
	sort.Ints(resArr)
	foot := &ListNode{}
	head := foot
	for i := 0; i< len(resArr); i++ {
		foot.Next = &ListNode{
			Val: resArr[i],
		}
		foot = foot.Next
	}
	return head.Next
}
