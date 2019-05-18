package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	l := len(flowerbed)
	for i := 0; i < l; i++ {
		if flowerbed[i] == 1 {
			continue
		}

		if (i+1 < len(flowerbed) && flowerbed[i+1] == 0 || i+1 == len(flowerbed)) &&
			(i-1 >= 0 && flowerbed[i-1] == 0 || i == 0) {
			flowerbed[i] = 1
			n--
		}
		if n == 0 {
			return true
		}
	}

	return false
}

func main() {
	input := []int{0, 1, 0, 1, 0, 0}
	n := 1
	fmt.Println(canPlaceFlowers(input, n))
}
