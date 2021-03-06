package main

import (
	"fmt"
)

func hammingDistance(x int, y int) int {
	var distance int
	val := x ^ y
	for {
		if val == 0 {
			break
		}
		distance++
		val &= val - 1
	}

	return distance
}

func main() {
	x := 1
	y := 4
	fmt.Println(hammingDistance(x, y))
}
