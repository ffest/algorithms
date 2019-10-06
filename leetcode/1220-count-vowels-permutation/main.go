package main

import (
	"fmt"
)

func countVowelPermutation(n int) int {
	var mod int64 = 1000000007
	cache := [5]int64{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		cacheCopy := [5]int64{}
		cacheCopy[0] += cache[1] % 1000000007
		cacheCopy[1] += (cache[0] + cache[2]) % 1000000007
		cacheCopy[2] += (cache[0] + cache[1] + cache[3] + cache[4]) % 1000000007
		cacheCopy[3] += (cache[2] + cache[4]) % 1000000007
		cacheCopy[4] += cache[0] % 1000000007
		copy(cache[:], cacheCopy[:])
	}

	var result int64
	for _, counter := range cache {
		result += counter
	}
	return int(result % mod)
}

// Slow recursive solution (easy to return all the combinations)
/*func countVowelPermutation(n int) int {
	if n == 0 {
		return 0
	}
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	var counter int
	helper('#', vowels, n, 0, &counter)
	return counter
}

func helper(subres byte, vowels []byte, n, level int, counter *int) {
	if level == n {
		*counter++
		*counter = *counter % 1000000007
		return
	}

	for _, vowel := range vowels {
		if subres == '#' {
			helper(vowel, vowels, n, level+1, counter)
			continue
		}
		switch {
		case subres == 'a' && vowel != 'e':
			continue
		case subres == 'e' && vowel != 'a' && vowel != 'i':
			continue
		case subres == 'i' && vowel == 'i':
			continue
		case subres == 'o' && vowel != 'i' && vowel != 'u':
			continue
		case subres == 'u' && vowel != 'a':
			continue
		}
		helper(vowel, vowels, n, level+1, counter)
	}
}*/

func main() {
	fmt.Println(countVowelPermutation(2))
}
