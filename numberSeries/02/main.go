package main

import "fmt"

//2. WAP to find a Number is Even or Odd

func main() {
	EvenOrOdd(12)
}

func EvenOrOdd(num int) {
	if num%2 == 0 {
		fmt.Printf("%d is even number", num)
	} else {
		fmt.Printf("%d is odd number", num)
	}
}
