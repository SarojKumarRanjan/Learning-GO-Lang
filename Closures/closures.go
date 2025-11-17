package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
		count += 1

		return count
	}

}

func main() {

	incr := counter()

	fmt.Println(incr())
	fmt.Println(incr())

	// the above thing will return 1 and 2
	// but how is that possible that it rem the prev value of count
	// and on second call return the next counter
	// this happens because of the closures

}
