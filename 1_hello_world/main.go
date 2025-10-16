package main //group of files called package
//it means this file belongs to the main package — the entry point of our Go program.

import "fmt"

/*
We are bringing in another package — specifically, the fmt package from Go’s standard library.
it’s not “inside” our main package.
It’s a separate, pre-built package that provides functions for formatting and printing (like fmt.Println, fmt.Printf, etc.).
Think of it as saying:
“Hey, I want to use some functions defined in the fmt package.”
*/

func main() { //entry point of the go app
	fmt.Println("hello world") //fmt is the package of standard lib use for formatting , print is the method of that package
}
