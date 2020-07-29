package main

//第一遍循环,将i位置上的数字交换到nums[i] - 1的位置(排除小于0,大于长度的数)
//第二遍循环找到的第一个不在位置上的数,即为所求
//不限制空间的话可以用堆来做

func firstMissingPositive(nums []int) int {

	for i:=0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] < len(nums) && nums[nums[i] - 1] != nums[i] {
			nums[nums[i] - 1], nums[i] = nums[i], nums[nums[i] - 1]
		}
	}

	for i:=0; i < len(nums); i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

