package main

import "fmt"

func main() {
	//var name = "Rajesh Kataraki"
	var age int32 = 37
	const isCool = true
	//shorthand
	name, email := "Rajesh", "rajeshkataraki@gmail.com"
//	isCool = false

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", name)
}
