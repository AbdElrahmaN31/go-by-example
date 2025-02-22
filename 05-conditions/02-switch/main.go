package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Not found")
	}

	switch time.Now().Weekday() {
	case time.Friday, time.Saturday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's morning")
	case t.Hour() < 17:
		fmt.Println("It's after noon")
	default:
		fmt.Println("It's evening")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an Integer")
		case bool:
			fmt.Println("It's a boolean")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(10)
	whatAmI(false)
	whatAmI("Hey!")
}
