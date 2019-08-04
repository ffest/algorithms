package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mapping := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	result := make([]string, 0)
	combinations("", digits, 0, mapping, &result)
	return result
}

func combinations(current, digits string, level int, mapping map[byte][]byte, result *[]string) {
	if len(current) == len(digits) {
		*result = append(*result, current)
		return
	}
	for _, letter := range mapping[digits[level]] {
		combinations(current+string(letter), digits, level+1, mapping, result)
	}
}

func main() {
	input := "23"
	fmt.Println(letterCombinations(input))
}
