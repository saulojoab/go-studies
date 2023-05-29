package main

import "fmt"

func main(){
	var x string = "Hello world"
	var y int = 10
	var z float64 = 0.2

	inferredString := "this is a string"
	inferredInt := 10
	inferredFloat := 0.2

	multipleString, multipleInt, multipleFloat := "this is a string", 10, 0.2

	fmt.Println("explicit types")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("implicit types")
	fmt.Println(inferredString)
	fmt.Println(inferredInt)
	fmt.Println(inferredFloat)

	fmt.Println("multiple variables")
	fmt.Println(multipleString)
	fmt.Println(multipleInt)
	fmt.Println(multipleFloat)
}