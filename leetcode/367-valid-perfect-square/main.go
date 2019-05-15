package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	guess := num / 2

	guesses := make(map[int]struct{})
	for {
		guesses[guess] = struct{}{}
		if guess*guess == num {
			return true
		}
		guess = (guess + num/guess) / 2
		if _, ok := guesses[guess]; ok {
			return false
		}
	}

	return false
}

func main() {
	input := 16
	fmt.Println(isPerfectSquare(input))
}
