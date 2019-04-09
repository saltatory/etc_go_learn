package main

import "fmt"
import "time"

func main() {
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("Weekend")
		default:
			fmt.Println("Weekday")
	}

	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Morning")
		default:
			fmt.Println("Night")
	}
}
