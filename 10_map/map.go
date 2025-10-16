package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string) // name := make(map[keyType]valType)

	//setting an element
	m["name"] = "manish"
	m["area"] = "backend"

	//getting
	fmt.Println(m["name"], m["area"]) //manish backend
	fmt.Println(m["notInTheMap"])     //"" empty string -> if key not found in the map , if key type is int then we get 0

	//clear the map
	clear(m)
	fmt.Println(m) //map[]

	m1 := make(map[string]int)
	m1["age"] = 10
	fmt.Println(m1["age"])   //1
	fmt.Println(m1["phone"]) //0 because of int , if key will be bool so bool's default should be False
	m1["price"] = 30
	fmt.Println(m1) //map[age:1 price:30]

	//delete from map
	fmt.Println(m1) //map[age:1 price:30]
	delete(m1, "price")
	fmt.Println(m1) //map[age:1]

	//finding an element

	// this will return 2 things _,_ if element found first _ would be the value of the key othervise emptyString or 0, or F according to the datatype second _ give boolean T/F accordingly contain can be anything "ok","isPresent"
	val, contain := m1["age"]
	fmt.Println(val) // 10
	if contain {
		fmt.Println("found") //True
	} else {
		fmt.Println("not found")
	}

	fmt.Println("=======COMPARE=======")
	m2 := map[string]int{"age": 1, "price": 20}
	m3 := map[string]int{"age": 1, "price": 20}

	fmt.Println(maps.Equal(m2, m3)) //true

}
