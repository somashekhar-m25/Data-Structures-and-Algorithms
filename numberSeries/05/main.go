package main

import "fmt"

//5. WAP to swap 2 numbers with 3 variables
func main() {
	SwapTwoNumbersWith3Variables(3, 4)
}
func SwapTwoNumbersWith3Variables(num1, num2 int) {
	fmt.Printf("num1: %d and num2: %d\n", num1, num2)
	temp := num1
	num1 = num2
	num2 = temp
	fmt.Printf("num1: %d and num2: %d", num1, num2)
}
