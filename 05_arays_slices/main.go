package main

import "fmt"

func main() {
	// Arrays
	var frutArr [2]string

	// Assign valuse

	frutArr[0] = "Apple"
	frutArr[1] = "Orange"

	// Declare and assing
	vegArr := [2]string{"Green Beans", "Carrots"}

	fmt.Println(frutArr)
	fmt.Println(vegArr)

	// Slices
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[0:2])
}
