package main

import "fmt"

func main(){
	x := 10
	y := 10

	// If else
	if x <= y {
		fmt.Printf("%d is less than or equal %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//else if
	color := "orange"

	if color == "red" {
		fmt.Println("Color is red")
	}else if color == "blue" {
		fmt.Println("Color is blue")
	} else{
		fmt.Println("Color is NOT red or blue")
	}

	//Switch

	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is NOT red or blue")
	}
}