package main

import (
	"fmt"
	sort2 "sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	l := len(nums)
	sort2.Ints(nums)

	for i := 0; i < l-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		st, ed := i+1, l-1
		for st < ed {
			total := nums[i] + nums[st] + nums[ed]
			if total > 0 {
				ed--
			} else if total < 0 {
				st++
			} else {
				res = append(res, []int{nums[i], nums[st], nums[ed]})
				for st < ed && nums[st] == nums[st+1] {
					st++
				}
				for st < ed && nums[ed] == nums[ed-1] {
					ed--
				}
				st++
				ed--
			}
		}
	}

	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}
