package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("s[2]: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	fmt.Println("s append - 1: ", s)
	s = append(s, "e", "f")
	fmt.Println("s append - 2: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[2:]
	fmt.Println("sl3: ", l)

	twoDimensional := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDimensional[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDimensional[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoDimensional)
}
