package main

import (
	"fmt"
)

func lemonadeChange(bills []int) bool {
	billMap := make(map[int]int)

	for _, b := range bills {
		switch b {
		case 5:
			billMap[5]++
		case 10:
			fifthsLeft := billMap[5]
			if fifthsLeft == 0 {
				return false
			}
			billMap[5] = fifthsLeft - 1
			billMap[10]++
		case 20:
			tensLeft := billMap[10]
			fifthsLeft := billMap[5]
			if tensLeft > 0 && fifthsLeft > 0 {
				billMap[5] = fifthsLeft - 1
				billMap[10] = tensLeft - 1
				continue
			} else if tensLeft == 0 && fifthsLeft >= 3 {
				billMap[5] = fifthsLeft - 3
				continue
			}
			return false
		}
	}

	return true
}

func main() {
	bills := []int{5, 10, 20}
	fmt.Println(lemonadeChange(bills))
}
