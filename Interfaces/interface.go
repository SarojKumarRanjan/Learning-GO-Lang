package main

import "fmt"

// interface is used to avoid open close from solid principle
// i.e the struct should be open to extension and close for modification

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {

	p.gateway.pay(amount)

}

// type razorPay struct{}

// func (p razorPay) pay(amount float32) {

// 	fmt.Println("Making paymnet using razorpay", amount)

// }

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment with stripe", amount)
}

func main() {

	// make payment
	//  gateway := razorPay{}
	gateway := stripe{}

	myPayment := payment{
		gateway: gateway,
	}

	myPayment.makePayment(34)

}
