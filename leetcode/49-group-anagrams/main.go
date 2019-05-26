package main

import "fmt"

func groupAnagrams(words []string) [][]string {
	letters := []byte("abcdefghijklmnopqrstuvwxyz")
	lscores := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	scores := make(map[byte]uint64, len(letters))
	for i, l := range letters {
		scores[l] = lscores[i]
	}

	groups := make(map[uint64][]string)
	for _, word := range words {
		ws := uint64(1)
		for _, b := range []byte(word) {
			ws *= scores[b]
		}
		groups[ws] = append(groups[ws], word)
	}
	r := make([][]string, 0)
	for _, g := range groups {
		r = append(r, g)
	}
	return r
}

func main() {
	fmt.Println(groupAnagrams([]string{"perfect", "wreathe", "listen", "inlets", "weather", "prefect", "enlist"}))
	fmt.Println(groupAnagrams([]string{"aa", "b"}))
}
