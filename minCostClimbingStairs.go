package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	n := len(cost) + 1
	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i:=2; i < n; i++ {
		_, min := sortTwo(dp[i-1], dp[i-2])
		if i == len(cost) {
			dp[i] = min
		} else {
			dp[i] = min + cost[i]
		}
	}
	return dp[n - 1]
}

func sortTwo(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}


func main() {
	a := []int{1, 100, 1, 1, 1, 1, 5}

	fmt.Println(minCostClimbingStairs(a))
}