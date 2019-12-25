package main

import "fmt"

func distributeCandies(candies int, num_people int) []int {
	distribution := make([]int, num_people)
	for i := 0; candies > 0; i++ {
		distribution[i%num_people] += min(candies, i+1)
		candies -= i + 1
	}
	return distribution
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	candies := 10
	people := 3
	fmt.Println(distributeCandies(candies, people))
}
