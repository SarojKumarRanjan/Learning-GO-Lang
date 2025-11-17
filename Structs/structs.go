package main

import (
	"fmt"
	"time"
)

// create a struct to represent a order of e-com

type order struct {
	id        string
	amount    float32
	status    bool
	createdAt time.Time // comes with nanosecond precision
}

// simulating a constructor function in go

func NewOrder(id string, amount float32, status bool) *order {

	// initial setup goes here
	newOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &newOrder
}

/* type transaction struct{
	id string
	name string
	order order

} */

// attaching a method to a struct

// here we will only use astrisk whenever we have to modify the value
func (o *order) changeStatus(status bool) {
	// here we do not need to de-refrence it as struct does for us
	// automatically
	o.status = status

}

func (o *order) getAmount() float32 {
	return o.amount
}

func main() {

	order := NewOrder("123", 23.34, true)

	order.createdAt = time.Now()
	order.changeStatus(true)

	fmt.Println(order.getAmount(), "this is the order amount")

	fmt.Println(order)

	// we can do this as well

	lang := struct {
		name string
	}{
		"golang",
	}
	fmt.Println(lang)
}
