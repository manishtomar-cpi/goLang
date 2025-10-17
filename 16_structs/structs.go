package main

import (
	"fmt"
	"time"
)

/*
Structs basically the custom data structures like classes, in go we dont have classes
when we want to group multiple fileds we can use structs
Works as blueprint
*/

type order struct { //making struct of order syntax is: type struct_name struct
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond
}

// reciever method type -> how to relate the methods in structs
func (o *order) changeStatus(status string) { //if we dont pass by ref it will not change -> only update when we pass by ref
	o.status = status //updating the status, struct doing deref automatically
}

// getter
func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	//making instance, if we are not setting the val of any field by default that will be the zero val of the type i.e int = 0

	myOrder := order{
		id:     "1",
		amount: 45.00,
		status: "recieved",
		//no need to pass all filds
	}
	//we can access the fields by '.' like javascript
	myOrder.createdAt = time.Now()

	fmt.Println(myOrder)

	//changeStatus("shipped") // give error ndefined: changeStatus
	myOrder.changeStatus("shipped")
	fmt.Println("amount of the order is: ", myOrder.getAmount())

	fmt.Println("after updated status", myOrder)

	firstStudent := newStudent("1", "manish", "10th")
	fmt.Println(firstStudent)
	fmt.Println(firstStudent.name)

	//inline structs used of declearing something
	language := struct {
		name   string
		isGood bool
	}{"goLang", true}

	fmt.Println(language)
}
