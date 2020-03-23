package main

import (
	"fmt"
	"strings"
)

func alienOrder(words []string) string {
	if len(words) == 0 {
		return ""
	}
	var result string
	cache := make(map[byte]map[byte]struct{})
	degrees := make(map[byte]int)
	for _, w := range words {
		for i := range w {
			degrees[w[i]] = 0
			if _, ok := cache[w[i]]; !ok {
				cache[w[i]] = make(map[byte]struct{})
			}
		}
	}

	for i := 0; i < len(words)-1; i++ {
		current, next := words[i], words[i+1]
		if len(current) > len(next) && strings.HasPrefix(current, next) {
			return ""
		}
		length := min(len(current), len(next))
		for j := 0; j < length; j++ {
			c1, c2 := current[j], next[j]
			if c1 != c2 {
				if _, ok := cache[c1][c2]; !ok {
					cache[c1][c2] = struct{}{}
					degrees[c2]++
				}
				break
			}
		}
	}
	queue := make([]byte, 0)
	for c, degree := range degrees {
		if degree == 0 {
			queue = append(queue, c)
		}
	}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result += string(current)
		for c := range cache[current] {
			degrees[c]--
			if degrees[c] == 0 {
				queue = append(queue, c)
			}
		}
	}
	if len(result) != len(degrees) {
		return ""
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	words := []string{"wrt", "wrf", "er", "ett", "rftt", "te"}
	fmt.Println(alienOrder(words))
}
