package main

import "fmt"

// enumerated type
type OrderStatus int

const (
	Recieved  OrderStatus = iota //0 indexed - it increments automatically for each new line within the same const block.
	Confrimed                    //1
	Prepared                     //2
	Shipped                      //3
	Delivered                    //4
)

func changeStatus(status OrderStatus) {
	fmt.Println("updated order status to", status)
}
func main() {
	changeStatus(Shipped)
}
