package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		if target == nums[0] {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}
    lo := findFirst(nums, target)
    if lo == -1 {
    	return []int{-1, -1}
	}
    hi := findLast(nums, target, lo)
	return []int{lo, hi}
}


func findFirst(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1

	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if nums[lo] == target {
		return lo
	} else {
		return -1
	}
}

func findLast(nums []int, target, lo int) int {
	hi := len(nums) - 1

	for lo < hi - 1 {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	if nums[hi] == target {
		return hi
	} else if nums[lo] == target {
		return lo
	}else {
		return -1
	}
}

func main() {
	//nums := []int{2, 2}
	nums := []int{5,6,7,8,8,10}
	fmt.Println(searchRange(nums, 8))
}