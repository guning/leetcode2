package list

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
}

func GenerateList(nums []int) *ListNode {
	var head, foot *ListNode
	foot = &ListNode{}
	head = foot
	for _, v := range nums {
		tmp := &ListNode{
			Val: v,
		}
		foot.Next = tmp
		foot = tmp
	}
	return head.Next
}

func PrintList(head *ListNode) {
	node := head
	for node != nil {
		fmt.Print(node.Val, "->")
		node = node.Next
	}
	fmt.Println("")
}

