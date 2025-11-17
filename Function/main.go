package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// function can return multiple values
func getLang() (int, string, int, string) {

	return 1, "hindi", 2, "english"
}

// in golang we can pass function in a function and also return it

func processIt(fn func(a int) string) {

	some := fn(2)
	fmt.Println(some)
}

func returnf() func(a int) int {

	return func(a int) int {
		return a
	}
}

func main() {

	result, b := add(2, 4), sub(3, 1)

	fmt.Println(result, b)

	a1, a2, a3, a4 := getLang()

	fmt.Println(a1, a2, a3, a4)

	fn := func(a int) string {

		if a == 5 {
			return "five"
		}

		return "not five"
	}

	processIt(fn)

	somef := returnf()

	f := somef(4)

	fmt.Println(f)

}
