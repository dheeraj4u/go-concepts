package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	para := strings.Fields(s)
	counter := make(map[string]int)
	for _, word := range para {
		counter[word]++
	}
	return counter
}

func main() {
	wc.Test(WordCount)
}
