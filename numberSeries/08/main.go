package main

import "fmt"

//8. WAP to swap 2 numbers with only 2 variables
func main() {
	SwapTwoNumbers(23, -89)
}

func SwapTwoNumbers(num1, num2 int) {
	fmt.Printf("num1 : %d and num2 : %d\n", num1, num2)
	// num1 = num1 + num2
	// num2 = num1 - num2
	// num1 = num1 - num2
	num1, num2 = num2, num1
	fmt.Printf("num1 : %d and num2 : %d\n", num1, num2)
}
