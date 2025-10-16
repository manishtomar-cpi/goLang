package main

import (
	"fmt"
	"time"
)

func simpleSwitch() {
	i := 4

	switch i {
	case 1:
		fmt.Println("one")
		//break -> no need to use the break in each case
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")

	}
}

func conditionalSwitch() {
	today := time.Now().Weekday()
	// fmt.Println(today)
	switch today {
	case time.Sunday, time.Saturday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Working day", today)
	}
}

// type switch
func whoAmI(i interface{}) { //interface means any type , i can be any type
	switch i.(type) {
	case int:
		fmt.Println("its a integer")
	case string:
		fmt.Println("its a string")
	case bool:
		fmt.Println("its a boolean")
	case float64:
		fmt.Println("its a float")
	default:
		fmt.Println("its a other")

	}
}
func main() {
	simpleSwitch()
	conditionalSwitch()
	whoAmI("hello")
}
