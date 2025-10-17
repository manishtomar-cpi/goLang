package main

import "fmt"

/*
Prevent the duplication -> generic
*/

func printSliceInt(nums []int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func printSliceString(nums []string) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

// now with generic
// T any means anytype can be ok, T int | string only int and string, same as template in C++
func printSlice[T int | string](nums []T) {
	//also can use "comparable" like func printSlice[T comparable](nums []T) -> comparable is an interface
	for _, num := range nums {
		fmt.Println(num)
	}
}

func main() {
	printSliceInt([]int{1, 2, 3}) //we need to make new function for all type for doing same operation (print here)
	printSliceString([]string{"manish", "ajay"})

	fmt.Println("+++++BY GENERIC+++++")

	nums := []int{1, 2, 3}
	names := []string{"manish", "sagar"}

	printSlice(nums)
	printSlice(names)
}
