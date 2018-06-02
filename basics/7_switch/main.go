package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Print("It's weekday")
	}

	// alternatif way to express if-else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's bedore noon")
	default:
		fmt.Println("It's afternoon")
	}

	// switch can also be used to compare variable types
	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Println("I don't know")
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("oi")
}
