package main

import "fmt"

// Declare constant
const Title string = "Person Details"
// const Title = "Person Details"

// Declare package variable
var Country string = "USA"
// var Country = "USA"

func main() {
	var fname, lname string = "John", "Doe"
	// fname, lname := "John", "Doe"
	var age int = 23
	// age := 23

	// Print constant variable
	fmt.Println(Title)
	// Print local variables
	fmt.Println("First Name:\t", fname)
	fmt.Println("Last Name:\t", lname)
	fmt.Println("Age:\t\t", age)
	fmt.Println("Country:\t", Country)
}