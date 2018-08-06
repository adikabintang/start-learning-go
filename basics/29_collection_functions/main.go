package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// like c++'s template class or java's generic
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}

	return true
}

func main() {
	var strs = []string{"peach", "pear", "plum", "apple"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "apple"))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
}
