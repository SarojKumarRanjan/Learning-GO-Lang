package main

import (
	"fmt"
	"slices"
)

func main() {
	// slices are the dynamic array of golang just like vectors in cpp
	// slices are the most used construct in go

	// this is uninitialized slice which is nil in go
	// var nums []int
	// fmt.Println(nums==nil)
	// fmt.Println(len(nums))

	// another way to initialize slices are

	// here the third params is the size of capacity if not provided
	// then automatically takes the size as its params
	// note very imp while appending if the elemnt cross the cap limit then go automatically
	// increase the capacity then the number you provided in the params as capacity

	nums := make([]int, 2, 5)

	// nums := []int{}

	// we can check the capacity of a slice

	fmt.Println(cap(nums))

	// This will give false as the size is 5 if the size is zero then its nil
	fmt.Println(nums == nil)

	// appending in slices

	nums = append(nums, 1, 2, 3, 4)
	// here the nums will be [0,1,2,3] it is init with zero and then
	//we are appending the next three value

	// now checking the capacity and size

	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	// slice operator

	fmt.Println(nums[0:3])

	nums2 := []int{6, 7, 8}

	// equal check in slices

	fmt.Println(slices.Equal(nums, nums2))

	/*
		--> there exist a copy function in slice
		--> slice operator to make any subarray
		--> we can make 2d slices as well nums := [][]int{{},{}}



	*/

}
