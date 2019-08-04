package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	var min, max string
	for _, str := range strs {
		if str == "" {
			return ""
		}
		if min == "" || min > str {
			min = str
		}
		if max == "" || max < str {
			max = str
		}
	}
	var result []byte
	for i := 0; i < len(min); i++ {
		if len(max) <= i || min[i] != max[i] {
			return string(result)
		}
		result = append(result, min[i])
	}

	return string(result)
}

func main() {
	input := []string{"flower", "flow", "flights"}
	fmt.Println(longestCommonPrefix(input))
}
