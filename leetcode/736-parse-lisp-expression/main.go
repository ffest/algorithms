package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func evaluate(expression string) int {
	return eval(expression, make(map[string]int))
}

func eval(expr string, parent map[string]int) int {
	if expr[0] != '(' {
		if expr[0] == '-' || (expr[0] >= '0' && expr[0] <= '9') {
			res, _ := strconv.Atoi(expr)
			return res
		}
		return parent[expr]
	}
	cache := make(map[string]int)
	for k, v := range parent {
		cache[k] = v
	}

	var tokens []string
	if expr[1] == 'm' {
		tokens = parse(expr[6 : len(expr)-1])
	} else {
		tokens = parse(expr[5 : len(expr)-1])
	}
	log.Print(expr, tokens)

	if strings.HasPrefix(expr, "(a") {
		return eval(tokens[0], cache) + eval(tokens[1], cache)
	} else if strings.HasPrefix(expr, "(m") {
		return eval(tokens[0], cache) * eval(tokens[1], cache)
	} else {
		for i := 0; i < len(tokens)-2; i += 2 {
			cache[tokens[i]] = eval(tokens[i+1], cache)
		}
		return eval(tokens[len(tokens)-1], cache)
	}
}

func parse(str string) []string {
	res := make([]string, 0)
	var parent int
	sb := make([]byte, 0)
	for i := range str {
		if str[i] == '(' {
			parent++
		}
		if str[i] == ')' {
			parent--
		}
		if parent == 0 && str[i] == ' ' {
			res = append(res, string(sb))
			sb = make([]byte, 0)
		} else {
			sb = append(sb, str[i])
		}
	}
	if len(sb) > 0 {
		res = append(res, string(sb))
	}
	return res
}

func main() {
	expression := "(let x 2 (mult x (let x 3 y 4 (add x y))))"
	fmt.Println(evaluate(expression))
}
