package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	dict := make([]int, 256)
	bs := []byte(s)
	var res, start int
	for i:=1; i<=len(s); i++ {
		if dict[int(bs[i-1])] > start {
			start = dict[int(bs[i-1])]
		}
		dict[int(bs[i-1])] = i
		if res < i - start {
			res = i - start
		}
	}
	return res

}
/*func lengthOfLongestSubstring(s string) int {
	res := 0
	bs := []byte(s)
	n := len(bs)
	for i := 0; i < n; i++ {
		tmp := make(map[byte]bool, n-i)
		//remain length<res, no need to calculate
		if n-i <= res {
			break
		}
		cnt := 0
		for j := i; j < n; j++ {
			if _, ok := tmp[bs[j]]; ok {
				break
			} else {
				tmp[bs[j]] = true
				cnt++
			}
		}
		if res < cnt {
			res = cnt
		}
	}

	return res
}*/

func main() {
	s := "au"
	fmt.Println(lengthOfLongestSubstring(s))
}
