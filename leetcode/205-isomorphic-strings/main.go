package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cache1 := make(map[byte]byte)
	cache2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, ok := cache1[s[i]]; ok && t[i] != v {
			cache1[s[i]] = t[i]
		}
		if v, ok := cache2[t[i]]; ok && s[i] != v {
			cache2[t[i]] = s[i]
		}
		cache1[s[i]] = t[i]
		cache2[t[i]] = s[i]
	}
	return true
}

func main() {
	s := "paper"
	t := "title"
	fmt.Println(isIsomorphic(s, t))
}
