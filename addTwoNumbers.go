package main

import (
	"fmt"
	"math/rand"
	"time"
	. "leetcode/list"
)
//Definition for singly-linked list.


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{
		Val: 0,
		Next:nil,
	}
	last := l3

	var add = 0
	for {
		if l1 == nil && l2 == nil {
			break
		}
		node := &ListNode {
			Val: 0,
			Next: nil,
		}
		var a,b int
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		res, nadd := addTN(a, b, add)
		add = nadd
		node.Val = res
		last.Next = node
		last = node
	}
	if add != 0 {
		node := &ListNode{
			Val: 1,
			Next:nil,
		}
		last.Next = node
	}
	return l3.Next
}

func addTN(a, b, c int) (int, int) {
	res := a + b + c
	if res >= 10 {
		return res%10, 1
	}
	return res, 0
}

func generateList(l *ListNode, n int) {
	last := l
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	last.Val = r.Intn(9)
	for i:=1; i<n; i++ {
		node := &ListNode {
			Val: 0,
			Next: nil,
		}
        node.Val = r.Intn(9)
        last.Next = node
        last = node
	}
}

func printList(l *ListNode) {
	for {
		fmt.Printf("%d ->", l.Val)
		if l.Next == nil {
			break
		}
		l = l.Next
	}
	fmt.Println("")
}


func main() {
    l1 := &ListNode{
    	Val: 0,
    	Next:nil,
	}
	l2 := &ListNode{
		Val: 0,
		Next:nil,
	}

    generateList(l1, 1)
    printList(l1)
	time.Sleep(2)
	generateList(l2, 2)
	printList(l2)

	l3 := addTwoNumbers(l1, l2)
	printList(l3)

}


