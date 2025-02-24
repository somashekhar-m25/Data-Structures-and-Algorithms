package main

import "fmt"

func main() {
	ArmstrongNumber(153)
}

func ArmstrongNumber(num int) {
	temp1 := num
	sum := 0
	for num > 0 {
		temp := num % 10
		sum = temp * temp * temp
		num = num / 10
	}
	if temp1 == sum {
		fmt.Printf("is Armstrong number: %d\n", temp1)
	} else {
		fmt.Printf("is Armstrong number: %d", temp1)
	}
}
