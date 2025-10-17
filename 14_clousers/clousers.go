package main

import "fmt"

func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}

/*
normally when we call the function it will come in call stack and the after execution the call stack of that function will clear but here as we can see we call counter first time the valeue is increased from 0 to 1 but remains in the call stack when we call again it is increasing to 2 because in the main the increment holds the count and it will remain in the stack while increment will be.

so in clouser what is happing if we are calling any variables from any outer scope that is remians insid the scope of that function

Same as javascript
*/
func main() {
	increment := counter()
	fmt.Println(increment()) //now increment is the function not variable which holds the count var
	fmt.Println(increment())

	/*
		increment holds the previous value of the outer scope (count). We call counter() once, so the returned function inside increment still remembers count, which is why calling increment() again increases it to 2.
	*/
}
