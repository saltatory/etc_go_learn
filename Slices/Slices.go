package main

import "fmt"
import "reflect"

func main() {
	// Simple slice
	a := make([]int,5)
	fmt.Println(a)

	// Array for funzies
	b:=[]int{1,2,3,4,5}
	fmt.Println("b",b)

	// Copy b into a
	copy(b,a)
	fmt.Println("b",b)
	b=[]int{1,2,3,4,5}

	// Print slice
	fmt.Println("slice",b[2:4])

	fmt.Println(reflect.TypeOf(b))
	x:=[5]int{1,2,3,4,5}
	fmt.Println(reflect.TypeOf(x))
}
