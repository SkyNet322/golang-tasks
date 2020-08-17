package golang

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	f := strings.Fields(s)
	for _, i := range f {
		m[i]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
