package main

import "fmt"

func totalFruit(tree []int) int {
	seen := [40001]int{}
	var count, current, left int
	for right := 0; right < len(tree); right++ {
		if seen[tree[right]] == 0 {
			current++
		}
		seen[tree[right]]++
		for current > 2 {
			seen[tree[left]]--
			if seen[tree[left]] == 0 {
				current--
			}
			left++
		}
		if count < right-left+1 {
			count = right - left + 1
		}
	}
	return count
}

func main() {
	tree := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fmt.Println(totalFruit(tree))
}
