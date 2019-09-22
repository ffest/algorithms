package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	firstRelease, secondRelease := 0, 0
	firstBuy, secondBuy := math.MaxInt32, math.MaxInt32

	for _, price := range prices {
		firstBuy = min(firstBuy, price)
		firstRelease = max(firstRelease, price-firstBuy)
		secondBuy = min(secondBuy, price-firstRelease)
		secondRelease = max(secondRelease, price-secondBuy)
	}

	return secondRelease
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
