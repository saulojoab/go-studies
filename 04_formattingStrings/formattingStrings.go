package main

import "fmt"

func main(){
	// %v is for when you want to print the value of a variable without specifying type
	// %s is for when you want to print a string
	// %d is for when you want to print an integer
	// %f is for when you want to print a float
	// %.2f is for when you want to print a float with 2 decimal places
	// %t is for when you want to print a boolean

	const name = "Saulo Joab"
	const age = 21
	const height = 1.75

	msg := fmt.Sprintf("Hi, my name is %s and I am %d years old. I also have a height of %.2f meters.", name, age, height)

	fmt.Println(msg)
}