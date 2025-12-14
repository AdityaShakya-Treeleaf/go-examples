package dsa

import "log"

// Count Covered Buildings
// You are given a positive integer n, representing an n x n city. You are also given a 2D grid buildings, where buildings[i] = [x, y] denotes a unique building located at coordinates [x, y].

// A building is covered if there is at least one building in all four directions: left, right, above, and below.

// Return the number of covered buildings.

// Example 1:

// Input: n = 3, buildings = [[1,2],[2,2],[3,2],[2,1],[2,3]]

// Output: 1

// Explanation:

// Only building [2,2] is covered as it has at least one building:
// above ([1,2])
// below ([3,2])
// left ([2,1])
// right ([2,3])
// Thus, the count of covered buildings is 1.
// Example 2:

// Input: n = 3, buildings = [[1,1],[1,2],[2,1],[2,2]]

// Output: 0

// Explanation:

// No building has at least one building in all four directions.
// Example 3:

// Input: n = 5, buildings = [[1,3],[3,2],[3,3],[3,5],[5,3]]

// Output: 1

// Explanation:

// Only building [3,3] is covered as it has at least one building:
// above ([1,3])
// below ([5,3])
// left ([3,2])
// right ([3,5])
// Thus, the count of covered buildings is 1.

type CountCoveredBuildings struct{}

func (c CountCoveredBuildings) CountCoveredBuildings(n int, buildings [][]int) int {
	if len(buildings) < 5 {
		return 0
	}

	covered := make([][]int, 0)
	for _, building := range buildings {
		if len(building) == 2 {
			x := building[0]
			y := building[1]

			if x-1 == 0 || y-1 == 0 || x+1 == n || y+1 == n {
				continue
			}

			covered = append(covered, []int{x, y})
		}
	}

	log.Printf("Covered Buildings: %d", covered)
	return len(covered)
}
