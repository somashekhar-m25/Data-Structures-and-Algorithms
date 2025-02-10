package main

import "fmt"

//11. Palindrome Number A palindrome number is a number that is the same after reverse. For example, 545, 151, 34543, 343, 171, 48984 are the palindrome numbers.
func main() {
	PalindromeNumber(1000002)
}

func PalindromeNumber(num int) {
	temp := num
	reverseNum := 0
	for num > 0 {
		reverseNum = reverseNum*10 + num%10
		num = num / 10
	}
	if reverseNum == temp {
		fmt.Printf("Given number %d is Plaindrome", temp)
	} else {
		fmt.Printf("Given number %d is not Plaindrome", temp)
	}
}
