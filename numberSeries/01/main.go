package main

import "fmt"

// Function to check if number is positive, negative, or zero
func PositiveOrNegative(num int) {
	if num > 0 {
		fmt.Printf("%d is postive number", num)
	} else if num < 0 {
		fmt.Printf("%d is negative number", num)
	} else {
		fmt.Printf("%d is neutral number", num)
	}
}

func main() {
	PositiveOrNegative(5)
}
