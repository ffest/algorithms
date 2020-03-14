package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	words := make(map[string]struct{})
	for _, word := range wordList {
		words[word] = struct{}{}
	}
	if _, ok := words[endWord]; !ok {
		return 0
	}
	queue := []string{beginWord}
	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}
	length := 1
	for len(queue) > 0 {
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			lastWord := queue[0]
			queue = queue[1:]
			for j := 0; j < len(lastWord); j++ {
				for k := 'a'; k <= 'z'; k++ {
					newWord := lastWord[:j] + string(k) + lastWord[j+1:]
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

	return 0
}

func main() {
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	beginWord := "hit"
	endWord := "cog"
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
