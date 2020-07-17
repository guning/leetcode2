package main

import (
	"fmt"
	"sort"
)

var res [][]int
func combinationSum(candidates []int, target int) [][]int {
	res = *new([][]int)
	sort.Ints(candidates)
	backtrack(target, candidates, []int{}, 0)

	return res
}

func backtrack(remain int, candidates []int, tmp []int, st int) {
	if remain < 0 {
		return
	}

	if remain == 0 {
		t := make([]int, len(tmp))
		copy(t, tmp)
		res = append(res, t)
		return
	}

	if remain > 0 {
		for i := st; i < len(candidates); i++ {
			t := tmp
			tmp = append(tmp, candidates[i])
			backtrack(remain-candidates[i], candidates, tmp, i)
			tmp = t
		}
	}
}

func main() {
	candicates := []int{2, 3, 5}
	res := combinationSum(candicates, 8)
	fmt.Println(res)
}
