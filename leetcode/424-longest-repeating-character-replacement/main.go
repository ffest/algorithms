package main

import "fmt"

// Check max possible length for every english uppercase letter
/*func characterReplacement(s string, k int) int {
	var maxLength int
	for i := 0; i < 26; i++ {
		ops := k
		start, end := 0, 0
		for end = 0; end < len(s); end++ {
			if s[end]-'A' != uint8(i) {
				for ops == 0 {
					if s[start]-'A' != uint8(i) {
						ops++
					}
					start++
				}
				ops--
			}
			if end-start+1 > maxLength {
				maxLength = end - start + 1
			}
		}
	}
	return maxLength
}*/

// Work with counts of all letters
func characterReplacement(s string, k int) int {
	var start, end, maxLength, maxCount int
	counts := [26]int{}
	for end = 0; end < len(s); end++ {
		counts[s[end]-'A']++
		if counts[s[end]-'A'] > maxCount {
			maxCount = counts[s[end]-'A']
		}
		// Window is too large
		for end-start+1 > k+maxCount {
			counts[s[start]-'A']--
			start++
		}
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}
	return maxLength
}

func main() {
	s := "AABABBA"
	k := 1
	fmt.Println(characterReplacement(s, k))
}
