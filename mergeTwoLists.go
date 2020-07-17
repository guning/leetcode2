package main

import . "leetcode/list"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var h, p *ListNode
	if l1.Val > l2.Val {
		h = l2
		p = l2
		l2 = l2.Next
	} else {
		h = l1
		p = l1
		l1 = l1.Next
	}
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			p.Next = l2
			p = p.Next
			l2 = l2.Next
		} else {
			p.Next = l1
			p = p.Next
			l1 = l1.Next
		}
	}

	if l1 != nil && l2 == nil {
		p.Next = l1
	} else {
		p.Next = l2
	}
	return h

}

func main() {
	l1 := GenerateList([]int{1,2,4})
	l2 := GenerateList([]int{1,3,4})
	PrintList(mergeTwoLists(l1, l2))
}

