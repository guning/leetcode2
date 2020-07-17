package main

import "strconv"

func myAtoi(str string) int {
	bs := []byte(str)
	isStart := true
	var tmp []byte
	for _, v := range bs {
		if v == ' ' && isStart {
			continue
		} else if (v == '-' || v == '+') && isStart {
			tmp = append(tmp, v)
			isStart = false
		} else {
			isStart = false
			if v >= '0' && v <= '9' {
				tmp = append(tmp, v)
			} else {
				break
			}
		}
	}
	res, _ := strconv.Atoi(string(tmp))
	if res < -1 << 31 {
		return -1 << 31
	}
	if res > 1 << 31 - 1 {
		return 1 << 31 - 1
	}
	return res
}