package main

import "fmt"

func ifElse() {
	age := 8

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else { //this else should be in the same line where if "}" end
		fmt.Println("Person is not and adult")
	}
}

func elseIf() {
	fmt.Println("====else if====")
	age := 10

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age > 12 {
		fmt.Println("Person is an teenager")
	} else {
		fmt.Println("Person is an kid")
	}
}

func logicalOperators() {
	role := "admin"
	hasPermissions := false

	if role == "admin" && hasPermissions { //not access because of false  if use || if can acess
		fmt.Println("you can access")
	}
}

// we can decleare a var inide the if constrants
func varInsideIfConstructs() {
	fmt.Println("=====varInsideIfConstructs=====")
	if age := 12; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Person is an teenager", age)
	} else {
		fmt.Println("Person is an kid", age)
	}
}

//go not have ternary opertor, we need to use if else

func main() {
	ifElse()
	elseIf()
	logicalOperators()
	varInsideIfConstructs()
}
