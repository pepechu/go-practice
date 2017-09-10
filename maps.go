package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	split := strings.Fields(s)
	m := make(map[string]int)
	for i := range split {
		v, b := m[split[i]]
		if b {
			m[split[i]] = v + 1
		} else {
			m[split[i]] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}