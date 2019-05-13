package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	var result []byte
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for _, str := range strs {
			if len(str) <= i || str[i] != c {
				return string(result)
			}
		}
	}
	return string(result)
}

func main() {
	input := []string{"flower", "flow", "flights"}
	fmt.Println(longestCommonPrefix(input))
}
