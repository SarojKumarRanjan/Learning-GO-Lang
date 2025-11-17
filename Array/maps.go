package main

import "fmt"

func main() {

	/*
		1. maps -> hashtables , object , dictionary




	*/

	//creating map

	// m := make(map[string]int)

	m := map[string]int{"some": 3}

	// setting an elemnts

	m["hello"] = 10
	m["mrng"] = 2

	// getting and element

	// if key is not present in map then we will get empty value
	fmt.Println(m["hello"])

	// we can get length by len() function
	fmt.Println(len(m))
	fmt.Println(m)

	// delete function in map

	// wen can also use clear function to empty map

	delete(m, "hello")

	fmt.Println(m)

	fp, ok := m["mrng"]

	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

	fmt.Println(fp)

}
