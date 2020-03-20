package main

import (
	"fmt"
	"sort"
)

type TrieNode struct {
	children [26]*TrieNode
	words    []string
}

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Slice(products, func(i, j int) bool {
		return products[i] < products[j]
	})
	root := &TrieNode{}
	for _, product := range products {
		node := root
		for _, c := range product {
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &TrieNode{
					words: make([]string, 0),
				}
			}
			node = node.children[c-'a']
			if len(node.words) < 3 {
				node.words = append(node.words, product)
			}
		}
	}

	result := make([][]string, len(searchWord))
	for i, c := range searchWord {
		if root == nil || root.children[c-'a'] == nil {
			root = nil
			result[i] = []string{}
			continue
		}
		root = root.children[c-'a']
		result[i] = root.words

	}
	return result
}

func main() {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	searchWord := "mouse"
	fmt.Println(suggestedProducts(products, searchWord))
}
