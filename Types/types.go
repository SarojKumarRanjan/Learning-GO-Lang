package main


import "fmt"


// short hand can not be used outside the function

// name := "hello go"

func main (){
	// fmt.Println("hello , printing string");
	// fmt.Println(1);

	// var name string  = "learning go lang"

	var name = "golang"

	//short hand syntax 
	// short hand can not be used outside the function

	name2:= "hello"

	fmt.Println(name,name2)

	//above thigs are for variables


	// now how to delare constant value in go 


	const something string = "hellp form someWhere"

	some := "what is this?"

	fmt.Println(something,some)


	//grouping the multiple constant 

	const (
		name1 = "helo"
		name3 = 5
		name4 = true
	)

	fmt.Println(name1,name3,name4)
}
