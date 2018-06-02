/*
Golang only have 'for' for looping
*/
package main

import (
	"fmt"
)

func main() {
	i := 1

	// it looks like 'while'
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// it looks like the classic for
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// it looks like while (1)
	k := 0
	for {
		fmt.Println("k: ", k)
		k++
		if k == 3 {
			break
		}
	}
}
