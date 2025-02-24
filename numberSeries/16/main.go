package main

import "fmt"

//16. WAP to Check Leap Year

func main() {
	leapyear(1900)
}

func leapyear(num int) {
	if num%4 == 0 && num%100 == 0 && num%400 == 0 {
		fmt.Printf("is leap year: %d", num)
	} else {
		fmt.Printf("is not leap year: %d", num)
	}
}
