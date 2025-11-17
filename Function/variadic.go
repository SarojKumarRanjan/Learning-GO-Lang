package main

import "fmt"

// something similar to rest and spread operator in js

// we can use interface {} for any type of params

// rest
func sum(nums ...int) int {
	addd := 0
	for i := range nums {
		addd += i
	}

	return addd
}

func main() {

	// variadic funtions a function which recives varibles number of parameters
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// spread
	s1 := sum(nums...)

	fmt.Println(s1)
}
