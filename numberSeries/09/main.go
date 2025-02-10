package main

import "fmt"

// 11. WAP to find given number is Buzz Number or not
// Numbers that are divisible by 7 or end with 7.
// For example, 57 is a buzz number because the number ends with seven.
// Another example is 28 because the number is divisible by seven.
func main() {
	BuzzNumber(21)
}

func BuzzNumber(num int) {
	if num%7 == 0 || num%10 == 7 {
		fmt.Printf("num %d is Buzz Number", num)
	} else {
		fmt.Printf("num %d is not Buzz Number", num)
	}
}
