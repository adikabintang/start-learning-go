package main

import (
	"fmt"
	"os"
)

func testingOne() {
	// defer works like stack: first in last out
	defer fmt.Println("test: A")

	fmt.Println("oi test nih")
	defer fmt.Println("test: B")
}

func createFile(p string) *os.File {
	fmt.Println("createFile")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("writeFile")
	fmt.Fprintln(f, "oi")
}

func closeFile(f *os.File) {
	f.Close()
}

func main() {
	testingOne()

	f := createFile("meme.txt")
	defer closeFile(f)
	writeFile(f)
}
