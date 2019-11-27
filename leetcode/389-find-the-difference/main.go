package main

import "fmt"

// XOR solution
func findTheDifference(s string, t string) byte {
	var diff byte
	for i := range s {
		diff ^= s[i]
		diff ^= t[i]
	}
	diff ^= t[len(t)-1]
	return diff
}

func main() {
	fmt.Println(string(findTheDifference("abcd", "abcde")))
}
