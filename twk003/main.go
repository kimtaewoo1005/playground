package main

import "fmt"

func main() {
	result := isValid("()[]{}")
	fmt.Println("Is valid:", result)
	return
}

// check if a string contains valid parentheses
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (c == ')' && top != '(') || (c == ']' && top != '[') || (c == '}' && top != '{') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
