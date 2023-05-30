package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wc := make(map[string]int)
	for _, w := range words {
		wc[w]++
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
