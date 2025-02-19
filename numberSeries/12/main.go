package main

import "fmt"

//12. WAP to Count Number of Digits in an Integer

func main() {
	CounterDigits(-12345678)
}

func CounterDigits(num int) {
	fmt.Printf("Given number : %d\n", num)
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}
	fmt.Printf("No.of digits in the given num is: %d", count)
}
