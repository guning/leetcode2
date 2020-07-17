package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := []string{""}
	dict := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	for i := range digits {
		var tmp []string
		index := digits[i] - '0'
		for _, v := range dict[index] {
			for _, vv := range res {
				tmp = append(tmp, vv + string(v))
			}
		}
		res = tmp
	}
	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
}