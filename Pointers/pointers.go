package main

import "fmt"

// recieve refrence
func somef(num *int) {

	// de-refrence
	*num = 5
	fmt.Println("inside func", num)

}

func main() {

	somet := 2

	// here we will send refrence
	somef(&somet)

	fmt.Println("in main func", somet)

}
