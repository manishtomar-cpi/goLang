package main

import "fmt"

// variadic functions are those functios in which we can pass "n" numbers of parameters like println
//make variadic function take any type use INTERFACE/any -> any type

func sum(nums ...int) int { // nums is the slice basically of int type
	total := 0
	for _, nums := range nums { //now we are iterating the slice to get the sum
		total += nums
	}
	return total
}

func main() {
	sum := sum(1, 2, 3, 4, 5)
	fmt.Println(sum)
}
