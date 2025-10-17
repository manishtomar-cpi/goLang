package main

import "fmt"

// we can use (a,b int) -> if all the parameters are of same type we can use the type at last
func add(a int, b int, c int) int {
	return a + b + c
}

// we can return multiple values in go, syntax is: func fun_name (params, type) (all return types){return in same type order}
func allOperationsInOneFunction(a, b int) (int, int, int, int) {
	return a + b, b - a, a * b, b / a
}

// passing function as a parameter
func processIt(fn func(a int) int) {
	t := fn(6)
	fmt.Println(t)
}

func main() {
	result := add(3, 5, 7)
	fmt.Println(result)

	sum, diffrence, multiplay, _ := allOperationsInOneFunction(2, 3) // if we dont want to use any return value we can use _ insted of this like we did in the case of divide above
	fmt.Println(sum, diffrence, multiplay)

	//define an anonymous function (a function without a name):
	fn := func(a int) int {
		return 3
	}
	processIt(fn) //return 3
}

/*

In Go, functions are first-class citizens — meaning we can:
	Assign them to variables
	Pass them as arguments
	Return them from other functions
*/
