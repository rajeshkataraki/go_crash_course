package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	//firstName string
	//lastName  string
	//city      string
	//gender    string
	//age       int
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old"
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried method (pointer receiver)
func (p *Person) getMarried(spouselastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouselastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Rajesh", lastName: "Kataraki", city: "Bangalore", gender: "m", age: 38}
	// Alternative
	person2 := Person{"Sneha", "Rao", "Delhi", "f", 35}

	//fmt.Println(person1)
	//fmt.Println(person1.firstName)
	//person1.age++
	//fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Patil")
	fmt.Println(person1.greet())
	person2.getMarried("Williams")
	fmt.Println(person2.greet())
}
