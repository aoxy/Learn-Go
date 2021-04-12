package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	wordList := strings.Fields(s)
	for _, w := range wordList {
		m[w] = m[w] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
