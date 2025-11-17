package main

import "fmt"

// range is used to iterate in data structure in go

func main() {

	nums := []int{1, 2, 3, 4, 5}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(i)
	// }

	ans := 0
	for v, i := range nums {

		//here v gives the index number number and i give the element on that index
		fmt.Println(v, "index")
		fmt.Println(i, "value ")
		ans += i
	}

	fmt.Println(ans)

	// iterating map using range

	obj := map[string]int{
		"hello": 2,
		"some1": 3,
	}

	for k, v := range obj {
		fmt.Println(k, v)
	}

	// unicode point rune
	//i is the starting point of rune
	for i, c := range "helo" {
		fmt.Println(i, c)
	}
}
