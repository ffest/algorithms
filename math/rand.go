package math

import (
	"time"
)

func Rand(lower, upper int) int {
	num := uint(upper - lower + 1)
	var result, i uint

	for {
		result, i = 0, 0
		for 1<<i < num {
			result = result<<1 | zeroOrOne()
			i++
		}
		if result <= num {
			break
		}
	}
	return int(result) + lower
}

func zeroOrOne() uint {
	return uint(time.Now().Unix() % 2)
}
