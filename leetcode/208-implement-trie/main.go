package main

import "fmt"

type Trie struct {
	isWord   bool
	children [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	current := t
	for _, c := range word {
		if current.children[c-'a'] == nil {
			current.children[c-'a'] = &Trie{}
		}
		current = current.children[c-'a']
	}
	current.isWord = true
}

func (t *Trie) Search(word string) bool {
	current := t
	for _, c := range word {
		if current.children[c-'a'] == nil {
			return false
		}
		current = current.children[c-'a']
	}
	return current.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t
	for _, c := range prefix {
		if current.children[c-'a'] == nil {
			return false
		}
		current = current.children[c-'a']
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
