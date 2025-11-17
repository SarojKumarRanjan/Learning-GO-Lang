package main

import "fmt"

// enums are enumerated types

type OrderStatus int

const (
	Recived OrderStatus = iota
	Delivered
	Packaged
	Confirmed
)

func checkStatus(s OrderStatus) {
	fmt.Println("changing status", s)
}

func main() {

	checkStatus(Recived)
}
