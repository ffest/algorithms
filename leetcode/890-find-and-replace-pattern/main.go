package main

import "fmt"

func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0, len(words))

	m := matchPattern(pattern)
	for _, s := range words {
		if matchPattern(s) == m {
			res = append(res, s)
		}
	}

	return res
}

func matchPattern(str string) int64 {
	var sum, i int64

	letterMap := make(map[rune]int64)
	for _, letter := range str {
		if letterMap[letter] == 0 {
			i++
			letterMap[letter] = i
		}
		sum = sum*10 + letterMap[letter]
	}
	return sum
}

func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"

	fmt.Println(findAndReplacePattern(words, pattern))
}
