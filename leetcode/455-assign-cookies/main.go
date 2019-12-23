package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var count int
	child, cookie := len(g)-1, len(s)-1
	for child >= 0 && cookie >= 0 {
		if g[child] <= s[cookie] {
			count++
			cookie--
		}
		child--
	}
	return count
}

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	fmt.Println(findContentChildren(g, s))
}
