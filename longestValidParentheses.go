package main

import (
	"fmt"
	. "leetcode/stack"
)

func longestValidParentheses(s string) int {
	stack := new(Stack)
	max := 0
	stack.Init(len(s))
    for i, v := range s {
    	if v == ')' {
    		if !stack.IsEmpty() {
    			if s[stack.Top()] == '(' {
    				stack.Pop()
				} else {
					stack.Push(i)
				}
			} else {
				stack.Push(i)
			}
		} else {
			stack.Push(i)
		}
	}

    if stack.IsEmpty() {
    	max = len(s)
	} else {
		//剩下的下标都是不匹配点的下标，以下计算不匹配点之间的距离，最长距离就是最终答案
		a := len(s)
		for !stack.IsEmpty() {
			b, _ := stack.Pop()
			max = getMax(max, a - b - 1)
			a = b
		}
		//0-a的长度
		max = getMax(a, max)
	}
    return max

}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	//fmt.Println(longestValidParentheses("()(()"))
	fmt.Println(longestValidParentheses(")()())"))
	//fmt.Println(longestValidParentheses("(()"))
}
