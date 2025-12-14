package dsa

//Number of Ways to Divide a Long Corridor
// Along a long library corridor, there is a line of seats and decorative plants. You are given a 0-indexed string corridor of length n consisting of letters 'S' and 'P' where each 'S' represents a seat and each 'P' represents a plant.

// One room divider has already been installed to the left of index 0, and another to the right of index n - 1. Additional room dividers can be installed. For each position between indices i - 1 and i (1 <= i <= n - 1), at most one divider can be installed.

// Divide the corridor into non-overlapping sections, where each section has exactly two seats with any number of plants. There may be multiple ways to perform the division. Two ways are different if there is a position with a room divider installed in the first way but not in the second way.

// Return the number of ways to divide the corridor. Since the answer may be very large, return it modulo 10^9 + 7. If there is no way, return 0.

// Example 1:

// Input: corridor = "SSPPSPS"
// Output: 3
// Explanation: There are 3 different ways to divide the corridor.
// The black bars in the above image indicate the two room dividers already installed.
// Note that in each of the ways, each section has exactly two seats.

type WaysDivideLongCoriddor struct{}

func (w WaysDivideLongCoriddor) NumberOfWays(corridor string) int {
	MOD := 1000000007

	seats := []int{}

	//find all positions of seats
	for i, c := range corridor {
		if c == 'S' {
			seats = append(seats, i)
		}
	}

	// if number of seats == 0, no dividers cannot be placed
	// if number of seats is not even, no section can have exactly 2 seats
	if len(seats) == 0 || len(seats)%2 != 0 {
		return 0
	}

	// if number of seats  == 2, only 1 divider can be placed
	if len(seats) == 2 {
		return 1
	}

	result := 1

	for i := 2; i < len(seats); i += 2 {
		// get gap between set of 2 seats
		gap := seats[i] - seats[i-1]

		// multiply result with each possibility i.e, placement in the gaps is a possibility
		result = (result * gap) % MOD
	}

	return result
}
