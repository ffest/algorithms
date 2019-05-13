package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var profit int
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > profit {
			profit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return profit
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(maxProfit(input))
}
