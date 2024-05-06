package main

import (
	"fmt"
	"math"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println("Minimum window substring:", minWindow(s, t))
}

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// Character count of t
	tCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}

	// Sliding window on s
	left, right, formed, required := 0, 0, 0, len(tCount)
	windowCounts := make(map[byte]int)
	minLen := math.MaxInt32
	minLeft := 0

	// Extend the window with the right pointer
	for right < len(s) {
		c := s[right]
		windowCounts[c]++

		// Only count characters that are in t
		if countInT, ok := tCount[c]; ok && windowCounts[c] == countInT {
			formed++
		}

		// Try and contract the window till the point it ceases to be 'desirable'.
		for left <= right && formed == required {
			c = s[left]

			// Save the smallest window that satisfies the condition
			if right-left+1 < minLen {
				minLen = right - left + 1
				minLeft = left
			}

			// The character at the position pointed by the `left` pointer is no longer a part of the window.
			windowCounts[c]--
			if countInT, ok := tCount[c]; ok && windowCounts[c] < countInT {
				formed--
			}

			// Move the left pointer ahead, this would help to look for a new window.
			left++
		}

		// Keep expanding the window once we are done contracting.
		right++
	}

	// Return the smallest window or an empty string if no such window exists.
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[minLeft : minLeft+minLen]
}
