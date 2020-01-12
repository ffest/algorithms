package main

import (
	"fmt"
)

func minFlips(a int, b int, c int) int {
	var flips int
	for c > 0 {
		if c&1 == 1 {
			if a&1 == 0 && b&1 == 0 {
				flips++
			}
		} else {
			if a&1 == 1 {
				flips++
			}
			if b&1 == 1 {
				flips++
			}
		}
		c >>= 1
		a >>= 1
		b >>= 1
	}
	for a > 0 {
		if a&1 == 1 {
			flips++
		}
		a >>= 1
	}
	for b > 0 {
		if b&1 == 1 {
			flips++
		}
		b >>= 1
	}
	return flips
}

func main() {
	a, b, c := 8, 3, 5
	fmt.Println(minFlips(a, b, c))
}
