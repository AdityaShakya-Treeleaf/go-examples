package dsa

import "log"

// Number of Smooth Descent Periods of a Stock
// You are given an integer array prices representing the daily price history of a stock, where prices[i] is the stock price on the ith day.

// A smooth descent period of a stock consists of one or more contiguous days such that the price on each day is lower than the price on the preceding day by exactly 1. The first day of the period is exempted from this rule.

// Return the number of smooth descent periods.

// Example 1:

// Input: prices = [3,2,1,4]
// Output: 7
// Explanation: There are 7 smooth descent periods:
// [3], [2], [1], [4], [3,2], [2,1], and [3,2,1]
// Note that a period with one day is a smooth descent period by the definition.
// Example 2:

// Input: prices = [8,6,7,7]
// Output: 4
// Explanation: There are 4 smooth descent periods: [8], [6], [7], and [7]
// Note that [8,6] is not a smooth descent period as 8 - 6 ≠ 1.

type SmoothDescentPeriods struct{}

func (s SmoothDescentPeriods) GetDescentPeriods(prices []int) int64 {
	n := len(prices)
	if n == 0 || n == 1 {
		return int64(n)
	}

	resultList := make([][]int, 0)
	for i, p := range prices {
		if i == 0 {
			resultList = append(resultList, []int{p})
			continue
		}

		dp := []int{p}
		resultList = append(resultList, dp)

		j := i
		curr := prices[j]
		prev := prices[j-1]
		log.Println(curr)
		log.Println(prev)
		for prices[j-1]-prices[j] == 1 {
			dp = append(dp, prev)
			resultList = append(resultList, dp)
			if j-1 == 0 {
				break
			} else {
				j--
			}
		}

		log.Println(resultList)
	}

	return int64(len(resultList))
}
