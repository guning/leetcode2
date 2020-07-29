package main


//https://www.bilibili.com/video/BV1NW411q7aM/
//https://www.youtube.com/watch?v=StH5vntauyQ
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	l, r := 0, len(height) - 1
	var res int
	lMax := height[l]
	rMax := height[r]

	for l < r {
		if lMax < rMax {
			res += lMax - height[l]
			l++
			if lMax < height[l] {
				lMax = height[l]
			}
		} else {
			res += rMax - height[r]
			r--
			if rMax < height[r] {
				rMax = height[r]
			}
		}
	}
	return res
}