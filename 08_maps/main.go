package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["Mike"] = "mike@gmail.com"
	emails["Fred"] = "fred@gmail.com"
	emails["Steph"] = "steph@gmail.com"

	// Declare map and add kv
	emails2 := map[string]string{"Mark": "mark@gmail.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map

	delete(emails, "Bob")
	fmt.Println(emails)

	fmt.Println(emails2)
}
