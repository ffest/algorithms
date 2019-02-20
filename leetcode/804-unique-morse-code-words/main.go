package main

import (
	"fmt"
)

func uniqueMorseRepresentations(words []string) int {
	morseCodes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	uniqMap := make(map[string]struct{})
	for _, word := range words {
		var morseWord string
		for _, b := range []byte(word) {
			morseWord += morseCodes[b-97]
		}
		uniqMap[morseWord] = struct{}{}
	}

	return len(uniqMap)
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}
