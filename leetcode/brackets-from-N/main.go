package main

import (
	"fmt"
	"log"
)

var usedLeft = 0
var usedRight = 0

func findAllCombinations(n int, current string) {
	log.Print(current, usedLeft, usedRight)
	if usedLeft == n && usedRight == n {
		fmt.Println(current)
		return
	}

	if usedLeft < n {
		current := current + "("
		usedLeft++
		findAllCombinations(n, current)
		if usedRight < usedLeft {
			current := current + ")"
			usedRight++
			findAllCombinations(n, current)
		}
	}
	if usedRight < usedLeft {
		current := current + ")"
		usedRight++
		findAllCombinations(n, current)
	}
}

func main() {
	var n = 3
	findAllCombinations(n, "")
}
