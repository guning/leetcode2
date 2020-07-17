package main

import (
	"fmt"
	"math/rand"
	"time"
)

func maxEqualRowsAfterFlips(matrix [][]int) int {
	max := 0
	for i := 0; i< len(matrix); i++ {
		tmp := 0
		for j:=i; j < len(matrix); j++ {
			if check(matrix[i], matrix[j]) {
				tmp++
			}
		}
		if max < tmp {
			max = tmp
		}
	}

	return max
}

func check(row1, row2 []int) bool {
	res := true
	first := row1[0] ^ row2[0]
	for i := 1; i < len(row1); i++ {
		tmp := row1[i] ^ row2[i]
		if tmp != first {
			res = false
			break
		}
	}
	return res
}
/*func maxEqualRowsAfterFlips0(matrix [][]int) int {
	max := 0
	colNums := len(matrix[0]) - 1
	nums := make([]int, len(matrix))
	OneList := generateOne(colNums+1)
	for i, v := range matrix {
		tmp := 0
		for j, vv := range v {
			tmp += int(math.Pow(2, float64(colNums - j))) * vv
		}
		nums[i] = tmp
	}

	for i := 0; i < len(nums); i++ {
		tmp := 0
		for j:=i; j < len(nums); j++ {
			res := nums[i] ^ nums[j]
			if res == 0 || res == OneList {
				tmp++
			}
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func generateOne(length int) int {
	var res float64
	for i:=length-1; i>=0; i-- {
		res += math.Pow(2, float64(i))
	}
	return int(res)
}*/

func main() {
	matric := make([][]int, 3)
	rand.Seed(time.Now().UnixNano())
	for i := range matric {
		matric[i] = make([]int, 3)
		for j := range matric[i] {
			matric[i][j] = rand.Intn(2)
		}
	}

	fmt.Println(matric)
	fmt.Println(maxEqualRowsAfterFlips(matric))
}
