package main


func reverse(x int) int {
	var res int
	var isNegative bool

	if x < 0 {
		isNegative = true
		x = -x
	}

	for x > 0 {
		remainder := x % 10
		x = (x - remainder) / 10
		res = res * 10 + remainder
	}

	if isNegative {
		res = -res
	}

	if res < -1 << 31 || res > 1 << 31 - 1 {
		return 0
	}

	return res
}