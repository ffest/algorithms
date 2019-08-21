package main

import (
	"fmt"
)

func countCharacters(words []string, chars string) int {
	if len(chars) == 0 {
		return 0
	}
	have := [26]int{}
	for _, c := range chars {
		have[c-'a']++
	}

	var sum int
	for _, w := range words {
		need := [26]int{}
		for _, c := range w {
			need[c-'a']++
		}

		goodString := true
		for i := 0; i < len(have); i++ {
			if need[i] > have[i] {
				goodString = false
				break
			}
		}
		if !goodString {
			continue
		}
		sum += len(w)
	}

	return sum
}

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	fmt.Println(countCharacters(words, chars))
}
