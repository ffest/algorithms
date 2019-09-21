package main

import "fmt"

func openLock(deadends []string, target string) int {
	deads := make(map[string]struct{})
	for _, dead := range deadends {
		deads[dead] = struct{}{}
	}
	visited := make(map[string]struct{})
	queue := []string{"0000"}
	var level int
	for len(queue) > 0 {
		qSize := len(queue)
		for i := 0; i < qSize; i++ {
			val := queue[0]
			queue = queue[1:]
			if _, ok := deads[val]; ok {
				continue
			}
			if val == target {
				return level
			}
			for j := 0; j < 4; j++ {
				next := getNextOrPrevValue(val[:j], val[j+1:], val[j], false)
				if _, ok := visited[next]; !ok {
					queue = append(queue, next)
					visited[next] = struct{}{}
				}
				prev := getNextOrPrevValue(val[:j], val[j+1:], val[j], true)
				if _, ok := visited[prev]; !ok {
					queue = append(queue, prev)
					visited[prev] = struct{}{}
				}
			}
		}
		level++
	}

	return -1
}

func getNextOrPrevValue(prefix, postfix string, char byte, prev bool) string {
	if prev {
		if char == '0' {
			char = '9'
		} else {
			char--
		}
	} else {
		if char == '9' {
			char = '0'
		} else {
			char++
		}
	}
	return prefix + string(char) + postfix
}

func main() {
	deadends := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	target := "8888"
	fmt.Println(openLock(deadends, target))
}
