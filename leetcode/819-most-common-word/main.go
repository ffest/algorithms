package main

import (
	"fmt"
)

func mostCommonWord(paragraph string, banned []string) string {
	bannedCache := make(map[string]struct{})
	for _, b := range banned {
		bannedCache[b] = struct{}{}
	}
	words := make([]string, 0)
	var tmp string
	for _, c := range paragraph {
		if c >= 'A' && c <= 'Z' {
			tmp += string(c + 32)
		} else if c >= 'a' && c <= 'z' {
			tmp += string(c)
		} else {
			if tmp != "" {
				words = append(words, tmp)
			}
			tmp = ""
		}
	}
	if tmp != "" {
		words = append(words, tmp)
	}

	cache := make(map[string]int)
	for _, w := range words {
		cache[w]++
	}
	max, maxWord := 0, ""
	for word, freq := range cache {
		if _, ok := bannedCache[word]; ok || (max > 0 && freq < max) {
			continue
		}
		max = freq
		maxWord = word
	}
	return maxWord
}

func main() {
	p := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	fmt.Println(mostCommonWord(p, banned))
}
