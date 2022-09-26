package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	r := make(map[string]int)
	ss := strings.Fields(s)
	for _, v := range ss {
		if r[v] == 0 {
			r[v] = 1
		} else {
			r[v] += 1
		}

	}
	return r
}

func main() {
	wc.Test(WordCount)
}
