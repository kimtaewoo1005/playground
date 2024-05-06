package main

import (
	"fmt"
)

func main() {
	n := int32(4) // Number of investments
	rounds := [][]int32{
		{2, 3, 603}, // round 1: contributions to investments 1 and 2
		{1, 1, 286}, // round 2: contributions to investments 3, 4, and 5
		{4, 4, 882}, // round 3: contributions to investments 2, 3, and 4
	}

	// Function call to find the maximum investment in any asset
	fmt.Println(maxValue(n, rounds))
}

func maxValue(n int32, rounds [][]int32) int64 {
	investments := make([]int64, n)

	for _, round := range rounds {
		left := round[0] - 1
		right := round[1] - 1
		contribution := int64(round[2])

		if left-1 < n {

		}
		investments[left] += contribution

		if right+1 < n {
			investments[right+1] -= contribution
		}
	}

	maxInvested := int64(0)
	currentInvestment := int64(0)

	for _, value := range investments {
		currentInvestment += value
		if currentInvestment > maxInvested {
			maxInvested = currentInvestment
		}
	}

	return maxInvested
}
