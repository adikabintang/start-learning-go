package main

import (
	"fmt"
	"sort"
)

// custom sort

type byLength []string

// three functions that have to be overloaded: Len(), Swap(), and Less()
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"mangga", "pisang", "rambutan"}

	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
