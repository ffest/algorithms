package main

import (
	"fmt"
)

type TrieNode struct {
	isWord bool
	next   [26]*TrieNode
}

type StreamChecker struct {
	root  *TrieNode
	query []byte
}

func Constructor(words []string) StreamChecker {
	root := &TrieNode{
		isWord: false,
		next:   [26]*TrieNode{},
	}
	for _, word := range words {
		current := root
		for i := len(word) - 1; i >= 0; i-- {
			if current.next[word[i]-'a'] == nil {
				current.next[word[i]-'a'] = &TrieNode{
					false,
					[26]*TrieNode{},
				}
			}
			current = current.next[word[i]-'a']
		}
		current.isWord = true
	}
	return StreamChecker{
		root:  root,
		query: make([]byte, 0),
	}
}

func (sc *StreamChecker) Query(letter byte) bool {
	sc.query = append(sc.query, letter)
	return sc.isWord()
}

func (sc *StreamChecker) isWord() bool {
	current := sc.root
	for i := len(sc.query) - 1; i >= 0; i-- {
		if current.next[sc.query[i]-'a'] == nil {
			return false
		}
		current = current.next[sc.query[i]-'a']
		if current.isWord {
			return true
		}
	}
	return false
}

func main() {
	sc := Constructor([]string{"abaa", "abaab", "aabbb", "bab", "ab"})
	fmt.Println(sc.Query('a'))
	fmt.Println(sc.Query('a'))
	fmt.Println(sc.Query('b'))
}
