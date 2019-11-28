package main

import (
	"fmt"
	"strings"
)

func frequencySort(s string) string {
	cache := make(map[byte]int)
	for i := range s {
		cache[s[i]]++
	}
	buckets := make([][]byte, len(s)+1)
	for b, f := range cache {
		buckets[f] = append(buckets[f], b)
	}

	var result string
	for i := len(buckets) - 1; i >= 0; i-- {
		bucket := buckets[i]
		if len(bucket) == 0 {
			continue
		}

		for j := 0; j < len(bucket); j++ {
			result += strings.Repeat(string(bucket[j]), i)
		}
	}
	return result
}

func main() {
	s := "tree"
	fmt.Println(frequencySort(s))
}
