package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	words := make(map[string]struct{})
	for _, w := range wordList {
		words[w] = struct{}{}
	}
	if _, ok := words[endWord]; !ok {
		return 0
	}
	var count int

	queue := []string{beginWord}
	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}
	length := 1
	for len(queue) > 0 {
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			lastWord := queue[0]
			queue = queue[1:]

			for i := 0; i < len(lastWord); i++ {
				for b := 'a'; b <= 'z'; b++ {
					newWord := lastWord[:i] + string(b) + lastWord[i+1:]
					if _, ok := words[newWord]; !ok || newWord == lastWord {
						continue
					}
					if newWord == endWord {
						return length + 1
					}
					if _, ok := visited[newWord]; !ok {
						queue = append(queue, newWord)
						visited[newWord] = struct{}{}
					}
				}
			}
		}
		length++
	}

	return count
}

func main() {
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	beginWord := "hit"
	endWord := "cog"
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
