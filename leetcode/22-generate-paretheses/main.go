package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	helper("", 0, 0, n, &result)
	return result
}

func helper(str string, opens, closes, n int, result *[]string) {
	if len(str) == 2*n {
		*result = append(*result, str)
		return
	}

	if opens < n {
		helper(str+"(", opens+1, closes, n, result)
	}
	if closes < opens {
		helper(str+")", opens, closes+1, n, result)
	}
}

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}
