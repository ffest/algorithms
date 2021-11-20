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

	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}
	queue := []string{beginWord}

	var length = 1
	for len(queue) > 0 {
		queueSize := len(queue)
		for k := 0; k < queueSize; k++ {
			lastWord := queue[0]
			queue = queue[1:]
			for i := 0; i < len(lastWord); i++ {
				for j := 'a'; j <= 'z'; j++ {
					newWord := lastWord[:i] + string(j) + lastWord[i+1:]
					if newWord == endWord {
						return length + 1
					}
					if _, ok := words[newWord]; !ok || newWord == lastWord {
						continue
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
