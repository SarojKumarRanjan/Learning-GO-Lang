package main

import "fmt"

/*
*
for loop is the only construct for looping in go lang
there does not exist while and do while loop
*/
func Main() {

	//simulating while loop using for loop

	// i := 1

	// for i <=5 {
	// 	//fmt.Println(i)
	// 	i++
	// }

	// for infinite loop
	// for i {
	// 	// some statment
	// }

	// now normal for loop

	for i := 0; i <= 5; i++ {

		// we can use break and continue keywords in this loop as in other languages

		fmt.Println(i)
	}

	// for go version 1.22 afterwards the new range thing is introduced

	//  for i:= 1 range 3{
	// 	fmt.Println(i)
	//  }

}
