package dsa

//Paths in Matrix Whose Sum Is Divisible by K
//You are given a 0-indexed m x n integer matrix grid and an integer k. You are currently at position (0, 0) and you want to reach position (m - 1, n - 1) moving only down or right.
//Return the number of paths where the sum of the elements on the path is divisible by k. Since the answer may be very large, return it modulo 109 + 7.
//
//Example 1:
//
//Input: grid = [[5,2,4],[3,0,5],[0,7,2]], k = 3
//Output: 2
//Explanation: There are two paths where the sum of the elements on the path is divisible by k.
//The first path highlighted in red has a sum of 5 + 2 + 4 + 5 + 2 = 18 which is divisible by 3.
//The second path highlighted in blue has a sum of 5 + 3 + 0 + 5 + 2 = 15 which is divisible by 3.

type PathsSumDivK struct{}

func (p PathsSumDivK) NumberOfPaths(grid [][]int, k int) int {
	MOD := 1000000007
	m, n := len(grid), len(grid[0])

	// dp[i][j][mod] = number of paths to (i,j) with sum%k = mod
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k)
		}
	}

	// Base case
	dp[0][0][grid[0][0]%k] = 1

	// Fill DP table
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}

			for prevMod := 0; prevMod < k; prevMod++ {
				newMod := (prevMod + grid[i][j]) % k

				// From top
				if i > 0 {
					dp[i][j][newMod] = (dp[i][j][newMod] + dp[i-1][j][prevMod]) % MOD
				}

				// From left
				if j > 0 {
					dp[i][j][newMod] = (dp[i][j][newMod] + dp[i][j-1][prevMod]) % MOD
				}
			}
		}
	}

	return dp[m-1][n-1][0]
}
