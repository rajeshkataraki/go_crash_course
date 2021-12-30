package main

import "fmt"

func greeting(name string) string{
	return "Hello " + name + "!!!"
}

func sum(num1, num2 int) int{
	return num1 + num2
}

func main(){
	fmt.Println(greeting("Rajesh"))
	fmt.Println(sum(7,3))
}