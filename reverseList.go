package main

import . "leetcode/list"

func reverseList(head *ListNode) *ListNode {
	newH := &ListNode{}
	for head != nil {
		tmp := head
		tmp.Next = newH
		newH = tmp
		head = head.Next
	}

	return newH
}

func main() {

}
