package main

import (
	"fmt"
)

func main() {

	// i := 2

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// 	// in go no need to write break statement here
	// case 2:
	// 	fmt.Println("two")
	// }

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its weekend")
	// default:
	// 	fmt.Println("its weekdays")
	// }

	// type switch

	WhoAmI := func(i interface{}) {

		switch t := i.(type) {
		case int:
			fmt.Println("this int ", t)
		case string:
			fmt.Println("this is string1", t)

		}

	}
	WhoAmI("hello")
	WhoAmI(1)

}
