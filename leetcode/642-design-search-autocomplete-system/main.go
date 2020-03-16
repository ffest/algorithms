package main

import (
	"fmt"
	"sort"
)

type counter struct {
	str     string
	counter int
}

type TrieNode struct {
	val      byte
	children map[byte]*TrieNode
	counters map[string]int
}

type AutocompleteSystem struct {
	node    *TrieNode
	search  string
	current *TrieNode
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	as := AutocompleteSystem{
		node: &TrieNode{
			children: make(map[byte]*TrieNode),
			counters: make(map[string]int),
		}}
	as.current = as.node
	for i, s := range sentences {
		as.add(s, times[i])
	}
	return as
}

func (as *AutocompleteSystem) add(sentence string, count int) {
	as.node.counters[sentence] = count
	node := as.node
	for i := range sentence {
		if node.children[sentence[i]] == nil {
			node.children[sentence[i]] = &TrieNode{
				val:      sentence[i],
				children: make(map[byte]*TrieNode),
				counters: make(map[string]int),
			}
		}
		node = node.children[sentence[i]]
		node.counters[sentence] = count
	}
}

func (as *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		as.node.counters[as.search]++
		as.add(as.search, as.node.counters[as.search])
		as.search = ""
		as.current = as.node
		return []string{}
	} else {
		as.search += string(c)
		if as.current == nil {
			return []string{}
		}
		as.current = as.current.children[c]
		if as.current == nil {
			return []string{}
		}
		result := make([]string, 0)

		counters := make([]counter, 0, len(as.current.counters))
		for k, v := range as.current.counters {
			counters = append(counters, counter{
				str:     k,
				counter: v,
			})
		}
		sort.Slice(counters, func(i, j int) bool {
			if counters[i].counter > counters[j].counter {
				return true
			} else if counters[i].counter < counters[j].counter {
				return false
			}
			return counters[i].str < counters[j].str
		})

		for i := 0; i < 3; i++ {
			if i == len(counters) {
				break
			}
			result = append(result, counters[i].str)
		}

		return result
	}
}

func main() {
	as := Constructor([]string{"i love you", "island", "ironman", "i love leetcode"}, []int{5, 3, 2, 2})
	fmt.Println(as.Input('i'))
	fmt.Println(as.Input(' '))
	fmt.Println(as.Input('a'))
	fmt.Println(as.Input('#'))
	fmt.Println(as.Input('i'))
	fmt.Println(as.Input(' '))
	fmt.Println(as.Input('a'))
	fmt.Println(as.Input('#'))
	fmt.Println(as.Input('i'))
	fmt.Println(as.Input(' '))
	fmt.Println(as.Input('a'))
	fmt.Println(as.Input('#'))
}
