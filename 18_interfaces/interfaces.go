package main

import "fmt"

// use to making our code scalable and organized
/*
assume we are making an payment app, and first we are using the razorpay know we want to go for stripe
so we need to make another struct as stripe and new pay function
then we need to go in the makePayment function and need to create another instance of strip and call the pay function from stripe
but we breaches the SOLID princle -> open close principle means our function is open for extend but close for modification
so what we can do we can add the field gateway in payment struct and pass the type which we want for now using stripe
then we just need to call p.gateway.pay(amount) and work done

but if we want to go back to razorpay again we need to update the payment struct type and the main function also accordingly
also it will give issue in the unit testing on make payment because make payment depends upon the gateway we need to intg test

interfaces works as a contract
we dont need to tell that "this" stuct follow interface, if the function signature inside the interface is same as inside the  struct function then go compilar understand that thic can use interface
we use "er" after variable when we make interface like "paymenter"
now we can do unit testing also because in testing function the function signature is same as in the interface
This enables polymorphism — writing code that works with different types as long as they satisfy the same interface.
everything is happing implicitly not expectily

so we are doing dependency invesion by using interface mean our main code is not depend on any concrit implementation
*/

type paymenter interface {
	pay(amount float32)
}
type payment struct {
	gateway paymenter //now the gateway of the paymenter type
}

// open close principle not following because we need to modify the code if we go from to stripe to razorpay
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGW := razorpay{}
	// razorpayPaymentGW.pay(amount)
	// 	stripePayment := stripe{}
	// 	stripePayment.pay(amount)

	//now we are following the open close rule because we just need to extend the payment struct to tell which one to use
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("payment processed by razorpay... ", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("payment processed by stripe...", amount)
}

// unit testing
type fakePayment struct{}

func (fp fakePayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway")
}
func main() {
	// newStripe := stripe{}
	fakePW := fakePayment{}
	newPayment := payment{
		gateway: fakePW,
	}
	newPayment.makePayment(100)
}
