package main

import "fmt"

func main() {

	// Using var
	// var name string = "Mike"
	var age int = 52
	const isCool = true

	// Shorthand
	// lastName := "Norton"

	name, lastName := "Liamwulf", "Cloudwalker"

	fmt.Println(name, lastName, age, isCool)
	fmt.Printf("%T\n", isCool)

}
