package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	ch := make(chan []int, 1)
	var res [][]int
	sort.Ints(candidates)
	go func() {
		backtrack(target, candidates, ch, []int{}, 0)
		close(ch)
	}()

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		res = append(res, v)
	}
	return res
}

func backtrack(remain int, candidates []int, ch chan []int, tmp []int, st int) {
	if remain < 0 {
		return
	}

	if remain == 0 {
		fmt.Println(remain, tmp)
		ch <- tmp
		return
	}

	if remain > 0 {
		for i := st; i < len(candidates); i++ {
			t := tmp
			tmp = append(tmp, candidates[i])
			backtrack(remain-candidates[i], candidates, ch, tmp, i)
			tmp = t
		}
	}
}

func main() {
	candicates := []int{2, 3, 5}
	res := combinationSum(candicates, 8)
	fmt.Println(res)
}
