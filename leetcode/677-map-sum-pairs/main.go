package main

import "log"

type MapSum struct {
	sum      int
	children map[byte]*MapSum
	original map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		children: make(map[byte]*MapSum),
		original: make(map[string]int),
	}
}

func (ms *MapSum) Insert(key string, val int) {
	diff := val
	if oldVal, ok := ms.original[key]; ok {
		diff -= oldVal
	}
	ms.original[key] = val

	current := ms
	for i := range key {
		if _, ok := current.children[key[i]]; !ok {
			trie := Constructor()
			current.children[key[i]] = &trie
		}
		current = current.children[key[i]]
		current.sum += diff
	}
}

func (ms *MapSum) Sum(prefix string) int {
	current := ms
	for i := range prefix {
		if child, ok := current.children[prefix[i]]; !ok {
			return 0
		} else {
			current = child
		}
	}
	return current.sum
}

func main() {
	mapSum := Constructor()
	mapSum.Insert("aa", 3)
	log.Print(mapSum.Sum("a"))
	mapSum.Insert("aa", 2)
	log.Print(mapSum.Sum("a"))
}
