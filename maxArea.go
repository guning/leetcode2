package main

func maxArea(height []int) int {
	var maxArea, st, ed int
	ed = len(height) - 1
	for st < ed {
		maxArea = max(maxArea, (ed - st) * min(height[st], height[ed]))
		if height[st] < height[ed] {
			st++
		} else {
			ed--
		}
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
