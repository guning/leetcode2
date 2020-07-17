package main

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i:=0; i< len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for j:=1; j < len(dp[0]); j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i:=1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			_, min := sortTwo(dp[i-1][j], dp[i][j-1])
			dp[i][j] = min + grid[i][j]
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func sortTwo(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

