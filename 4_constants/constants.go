package main

import "fmt"

// name:= "manish"  //sorthand decleration is not allowed outside the func we still can use const and other declerations
const name = "manish"

var age int = 26

func get_val() {
	fmt.Println(name, age)
}
func main() {
	const pi = 3.14
	const age int = 5 // both the ways is allowed to declare
	fmt.Println(pi, age)
	fmt.Println("getters")
	get_val()

	//const grouping
	const (
		port = 5000
		host = "local host"
	)
	fmt.Println(port, host)
}
