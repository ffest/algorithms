package main

import (
	"fmt"
	"log"
)

func findMaxPalindrom(str string) string {
	if len(str) <= 2 {
		return str[:1]
	}

	strBytes := []byte(str)

	var max []byte
	for i := range strBytes {
		if i == 0 || i == len(strBytes)-1 {
			continue
		}

		var left int
		var right int
		var localMax int
		left = i
		if strBytes[i] == strBytes[i+1] {
			right = i + 1
			localMax = 2
		} else {
			right = i
			localMax = 1
		}

		for {
			if left == 0 || right >= len(strBytes)-1 {
				break
			}

			left--
			right++
			if strBytes[left] != strBytes[right] {
				break
			}
			localMax += 2
		}

		if localMax > len(max) {
			log.Print(i)
			log.Print(string(strBytes[left : right+1]))
			max = strBytes[left : right+1]
		}
	}
	return string(max)
}

func main() {
	fmt.Println(findMaxPalindrom("ccaba"))
}
