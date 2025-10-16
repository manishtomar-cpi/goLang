package main

import "fmt"

func main() {
	fmt.Println("======range in slice======")
	arr := []int{1, 2, 3}

	//iterating the arr
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//iterting through range
	fmt.Println("=====iterting through range=====")
	sum := 0
	for i, num := range arr { //it will return two thigs _,_ first _ will be index and second _ will the val at that index
		sum += num
		fmt.Println(i, " -> ", num)
	}
	// 0  ->  1
	// 1  ->  2
	// 2  ->  3
	fmt.Println(sum, "is the sum of slice") //6

	fmt.Println("======range in map======")
	m := map[string]int{"price": 100, "unit": 2, "quantity": 10}
	for key, value := range m { //in map _,_ return key_value
		fmt.Println(key, " -> ", value) //price  ->  100, unit  ->  2, quantity  ->  10
	}
	for key := range m {
		fmt.Println(key) // price,unit, quantity - > only give keys now
	}
}
