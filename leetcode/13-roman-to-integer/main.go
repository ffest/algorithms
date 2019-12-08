package main

import "fmt"

func romanToInt(s string) int {
	symbols := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var sum int
	for i := len(s) - 1; i >= 0; i-- {
		current := symbols[s[i]]
		if i > 0 && symbols[s[i-1]] < current {
			current -= symbols[s[i-1]]
			i--
		}
		sum += current
	}
	return sum
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
