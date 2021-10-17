package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var prev_sell, sell, prev_buy, buy = 0, 0, 0, -math.MaxInt8
	for _, price := range prices {
		prev_buy = buy
		buy = max(prev_sell-price, prev_buy)
		prev_sell = sell
		sell = max(prev_buy+price, prev_sell)
	}
	return sell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var prices = []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}
