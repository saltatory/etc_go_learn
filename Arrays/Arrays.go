package main

import "fmt"

func main() {
	// Init
	var a5 [5]int

	// Print empty
	fmt.Println(a5)

	// Set one value
	a5[2] = 42
	fmt.Println(a5)
	fmt.Println("Length:",len(a5))

	// Static initializer
	var b5 [5]int
	b5 = [5]int{1,2,3,4,5}
	fmt.Println(b5)

	// Duck typer
	c5 := [5]int{1,2,3,4,5}
	fmt.Println(c5)

}
