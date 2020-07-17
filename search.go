package main

import "fmt"

func search(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1
	for lo <= hi{
		mid := (hi + lo) / 2
		if nums[mid] == target {
			return mid
		}

		//左边递增
		if nums[lo] <= nums[mid] {
			if target >= nums[lo] && target < nums[mid] {
				//target 在左边,右边都丢掉
				hi = mid - 1
			} else {
				//target 在右边，左边都丢掉
				lo = mid + 1
			}
		} else {
			//左边不递增右边肯定递增
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4,5,6,7,0,1,2}
	fmt.Println(search(nums, 0))
}