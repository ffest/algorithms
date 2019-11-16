package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	firstSell, secondSell := 0, 0
	firstBuy, secondBuy := math.MinInt32, math.MinInt32

	for _, price := range prices {
		secondSell = max(secondSell, secondBuy+price)
		secondBuy = max(secondBuy, firstSell-price)
		firstSell = max(firstSell, firstBuy+price)
		firstBuy = max(firstBuy, -price)
	}

	return secondSell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices))
}
