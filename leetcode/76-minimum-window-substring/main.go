package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	have := [128]int{}
	need := [128]int{}
	for i := range t {
		need[t[i]]++
	}

	result := ""
	var tail, head, count int
	for ; head < len(s); head++ {
		if have[s[head]] < need[s[head]] {
			count++
		}
		have[s[head]]++

		for tail <= head && have[s[tail]] > need[s[tail]] {
			have[s[tail]]--
			tail++
		}

		length := head - tail + 1
		if count == len(t) && (result == "" || len(result) > length) {
			result = s[tail : head+1]
		}
	}

	return result
}

func main() {
	s := "ab"
	t := "b"
	fmt.Println(minWindow(s, t))
}
