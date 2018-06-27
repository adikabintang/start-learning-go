package main

import (
	"fmt"
)

// Interfaces are named collections of method signatures
// An interface is two things: it is a set of methods, but it is also a type
// basic interface for geometric shapes, we’ll implement this interface on rectangle and circle types.
// method area and perim are generalized
type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// implement all methods in the interface
func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perim() float64 {
	return (r.width + r.height) * 2
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * 3.14 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rectangle{width: 3, height: 4}
	fmt.Printf("rectangle's area: %f\n", r.area())
	fmt.Printf("rectangle's perim: %f\n", r.perim())

	geo := []geometry{rectangle{width: 5, height: 6}, circle{radius: 7}}

	fmt.Printf("g (rectangle): %f\n", geo[0].area())
	fmt.Printf("g (circle): %f\n", geo[1].area())

	for _, g := range geo {
		measure(g)
	}
}
