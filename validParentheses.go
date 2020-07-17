package main

import "fmt"


type stack []int32
func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}
	if len(s) == 0 {
		return true
	}
	var stk stack
	stackin := func(b int32) {
		stk = append(stk, b)
	}

	stackgettop := func() int32 {
		return stk[len(stk) - 1]
	}

	stackout := func() {
		stk = stk[:len(stk)-1]
	}


	for _, v := range s {
		if len(stk) == 0 {
			if v != '(' && v != '[' && v != '{' {
				return false
			}
			stackin(v)
		} else {
			if v == '(' || v == '[' || v == '{' {
				stackin(v)
			} else {
				t := stackgettop()
				s1 := v == ')' && t == '('
				s2 := v == ']' && t == '['
				s3 := v == '}' && t == '{'
				if s1 || s2 || s3 {
					stackout()
					continue
				} else {
					return false
				}
			}
		}
	}

	if len(stk) != 0 {
		return false
	}

	return true
}

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}
