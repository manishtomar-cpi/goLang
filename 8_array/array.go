package main

import "fmt"

func Array() {
	var nums [5]int
	fmt.Println(len(nums)) //len gives us the len of the func

	nums[0] = 1
	fmt.Println(nums[0])

	fmt.Println(nums) // [1,0,0,0,0,0] -> we can print whole array only by println function

	//zeored values will change accoring to the type

	var values [4]bool
	values[1] = true
	fmt.Println(values) //[f,t,f,f,] -> by default all false, In case of string its all empty space by default

	var dec [4]float64
	dec[1] = 1.12
	fmt.Println(dec) //[0,1.12,0,0]

	//short hand dec of array
	numbers := [5]int{1, 2, 3, 4}
	fmt.Println(numbers) //[1,2,3,4,0]

}

func Array_2D() {
	nums := [2][2]int{{2, 3}}
	fmt.Println(nums) // [[2,3],[0,0]]
}

func main() {
	Array()
	Array_2D()
}
