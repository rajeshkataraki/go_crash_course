package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign key value
	//emails["Rajesh"] = "rajeshkataraki@gmail.com"
	// emails["Bob"] = "bob@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	emails := map[string]string{"Rajesh": "rajeshkataraki@gmail.com", "Bob": "bob@gmail.com"}
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete
	delete(emails, "Bib")
	fmt.Println(emails)
}
