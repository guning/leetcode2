package main

import (
	. "leetcode/list"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	var deleteNode, preNode, foot *ListNode
	var cnt int
	deleteNode = head
	foot = head

	for foot != nil {
		foot = foot.Next
		cnt++
		if cnt >= n  && foot != nil{
			preNode = deleteNode
			deleteNode = deleteNode.Next
		}
	}

	if deleteNode == head {
		return head.Next
	}
	preNode.Next = deleteNode.Next


	return head
}

func main() {
	head := GenerateList([]int{1,2,3,4,5})
	PrintList(head)
	PrintList(removeNthFromEnd(head, 2))
}
