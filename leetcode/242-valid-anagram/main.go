package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cache := [26]int{}
	for _, c := range s {
		cache[c-'a']++
	}

	for _, c := range t {
		if cache[c-'a'] == 0 {
			return false
		}
		cache[c-'a']--
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
