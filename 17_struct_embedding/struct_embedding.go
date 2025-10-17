package main

import "fmt"

type customer struct {
	id    string
	name  string
	phone string
}

type order struct {
	id       string
	amount   float32
	status   string
	customer //struct embedding
}

func main() {

	newOrder := order{
		id:     "1",
		amount: 50.09,
		status: "shipped",
	}

	fmt.Println(newOrder) // {1 50.09 shipped {  }} -> {   } is emoty string (Zero values) of string

	newCustomer := customer{
		id:    "1",
		name:  "manish",
		phone: "12121212",
	}
	newOrder.customer = newCustomer
	fmt.Println(newOrder) // {1 50.09 shipped {1 manish 12121212}} -> now taking the newCustomer

	newOrder.customer.name = "ajay" //means the order now have seprate customer and we can update the fields
	fmt.Println(newOrder)           // {1 50.09 shipped {1 ajay 12121212}} -> now taking the newCustomer
	fmt.Println(newCustomer)

}
