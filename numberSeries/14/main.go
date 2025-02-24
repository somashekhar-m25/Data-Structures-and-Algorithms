package main

import "fmt"

//14. Narcissistic Number is a number that is the sum of its own digits each raised to the power of the number of digits

func main() {
	NarcissisticNumber(1634)
}

func NarcissisticNumber(num int) {
	temp := num
	power := 0
	for num > 0 {
		power++
		num /= 10
	}
	num = temp
	sum := 0
	for num > 0 {
		sum += PowerOf(num%10, power)
		num /= 10
	}
	if temp == sum {
		fmt.Printf("is Narcissistic number: %d\n", temp)
	} else {
		fmt.Printf("is not Narcissistic number: %d\n", temp)
	}
}

func PowerOf(num, power int) int {
	v := 1
	for i := 1; i <= power; i++ {
		v *= num
	}
	return v
}
