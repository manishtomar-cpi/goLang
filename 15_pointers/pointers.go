package main

import "fmt"

func changeNum(num int) { //passing numm by value so the copy will generate
	fmt.Println("++++++BY VALUE++++++")
	num = 5
	println("change in chnageNum", num) //5
}

func changeNumByReference(num *int) {
	fmt.Println("++++++BY REF++++++")
	*num = 5 // we are first derefrencing the pointer by *num
}

func main() {
	num := 2
	changeNum(num)
	fmt.Println("after change by val: ", num) //1 -> because we passed by value
	// fmt.Println("memory of the num", &num)

	changeNumByReference(&num) //passing by ref

	fmt.Println("after change by ref: ", num) //5-> because we passed by value
}
