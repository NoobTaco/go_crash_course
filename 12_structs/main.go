package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getsMarried (Pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "Male" {
		return
	}
	p.lastName = spouseLastName

}

func main() {
	// init person using struct
	person1 := Person{
		firstName: "Mike", // First name
		lastName:  "Norton",
		city:      "Fresno",
		gender:    "Male",
		age:       52}

	person2 := Person{"Bethany", "Porter", "Tampa", "Female", 29}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	fmt.Println(person2)
	fmt.Println(person2.firstName)

	person1.hasBirthday()          // Adds to birthday.
	person2.getMarried("Williams") // Adds to birthday.
	fmt.Println(person2.greet())
}
