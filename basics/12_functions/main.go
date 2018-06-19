package main

import (
	"fmt"
)

func addition(a int, b int) int {
	return a + b
}

func aVoidFunc() {
	fmt.Println("a void function")
}

// multiple return function
func aFuncMultipleRet() (int, int) {
	return 4, 6
}

// variadic function
func sums(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

// closure: useful when you want to define a function inline without having to name it
func initSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	res := addition(12, 13)
	fmt.Println("res: ", res)

	aVoidFunc()
	a, b := aFuncMultipleRet()
	fmt.Printf("a: %d, b: %d\n", a, b)

	// ignore the first
	_, c := aFuncMultipleRet()
	fmt.Println("c: ", c)

	// use variadic
	sum := sums(1, 2, 3)
	fmt.Println("sums: ", sum)

	// use clouse
	nextInt := initSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := initSeq()
	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())

	// use recursive
	fmt.Println("factorial of 4: ", factorial(4))
}
