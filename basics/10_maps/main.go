package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["key1"] = 1
	m["key2"] = 2

	fmt.Println("map: ", m)

	var1 := m["key1"]
	fmt.Println("var1: ", var1)

	var1 = m["key2"]
	fmt.Println("var1: ", var1)

	fmt.Println("len: ", len(m))

	delete(m, "key2")
	fmt.Println("map: ", m)

	/*
		the "_, checkVar" notation is used to check whether the map has the value for that key
	*/
	_, isThisPresent := m["kancut"]
	if isThisPresent == false {
		fmt.Println("no value for kancut key")
	} else {
		fmt.Println("anotherVal: ", m["kancut"])
	}

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n: ", n)
}
