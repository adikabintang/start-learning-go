package main

import (
	"fmt"
	"math"
)

const CONST_VAR string = "oooi"

func main() {
	fmt.Println(CONST_VAR)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
