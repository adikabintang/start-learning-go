package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

// the struct is passed by value, the real struct is not changed
func (p person) sayItMan() {
	p.age++
	fmt.Println("age ", p.age)
}

// the struct is passed by pointer. the member can be changed
func (p *person) sayItManWithPointer() {
	p.age++
	fmt.Println("age ", p.age)
}

func (p person) getNumberPlusOne() int {
	return (p.age + 1)
}

func main() {
	fmt.Println(person{name: "Bob", age: 23})

	// An & prefix yields a pointer to the struct
	s := person{name: "Ibra", age: 40}
	sp := &s

	fmt.Println("s age: ", s.age)
	sp.age = 30
	fmt.Println("s age after: ", s.age)

	s.sayItMan()
	fmt.Println("s age: ", s.age)

	s.sayItManWithPointer()
	fmt.Println("s age: ", s.age)

	fmt.Println("getnumberplusone: ", s.getNumberPlusOne())
}
