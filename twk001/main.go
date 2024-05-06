package main

import "fmt"

func main() {
	result := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(result)
}

// return the length of the longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	// create a map to store the last index of each character
	lastIndex := make(map[rune]int)
	// initialize the start and max variables
	start, max := 0, 0
	// iterate over the string
	for i, c := range s {
		// if the character is already in the map
		if index, ok := lastIndex[c]; ok {
			// update the start variable to the next index
			start = maxInt(start, index+1)
		}
		// update the max variable
		max = maxInt(max, i-start+1)
		// update the index of the character
		lastIndex[c] = i
	}
	// return the max variable
	return max
}

// return the maximum of two integers
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// return the minimum of two integers
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
