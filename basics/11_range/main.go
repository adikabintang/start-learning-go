// range iterates over elements in a variety of data structures
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	sum := 0

	// use _, to ignore the index
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum = ", sum)

	for i, num := range nums {
		fmt.Printf("nums: %d, index %d\n", num, i)
	}

	aVarMap := map[string]string{"a": "nyo", "b": "nyi"}
	for key, val := range aVarMap {
		fmt.Printf("key: %s, value: %s\n", key, val)
	}

	// iterating over key only
	for key := range aVarMap {
		fmt.Println("key: ", key)
	}

	for index, character := range "abc" {
		fmt.Printf("index: %d, char: %c\n", index, character)
	}
}
