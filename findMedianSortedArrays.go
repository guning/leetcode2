package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	cursor1 := 0
	cursor2 := 0
	all := len(nums1) + len(nums2)
	if all == 1 {
		if len(nums1) == 1 {
			return float64(nums1[0])
		}
		if len(nums2) == 1 {
			return float64(nums2[0])
		}
	}
	isSingle := all%2 == 1
	k := all/2 + 1
	if isSingle {
		k = all/2 + 2
	}
	var a, b float64

	for i := 0; i < k; i++ {
		a = b
		//edge
		if cursor1 >= len(nums1) {
			b = float64(nums2[cursor2])
			cursor2++
			continue
		}
		if cursor2 >= len(nums2) {
			b = float64(nums1[cursor1])
			cursor1++
			continue
		}
		//normal
		if nums1[cursor1] <= nums2[cursor2] {
			b = float64(nums1[cursor1])
			cursor1++
		} else {
			b = float64(nums2[cursor2])
			cursor2++
		}

	}
	if isSingle {
		return a
	} else {
		return (a + b) / 2
	}
}

func main() {
	nums1 := []int{1,3}
	nums2 := []int{2,4,5}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
